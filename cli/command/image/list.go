package image

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/quarkloop/quarkloop/cli/client"
	"github.com/quarkloop/quarkloop/cli/console"
	"github.com/quarkloop/quarkloop/pkg/grpc/v1/cluster"
)

func NewImageListCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ls",
		Short: "List namespace images",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			namespace, _ := cmd.Flags().GetString("namespace")
			return runImageList(cmd.Context(), cmd, namespace)
		},
	}

	flags := cmd.Flags()
	flags.String("namespace", "", "image namespace")

	return cmd
}

type IterableImageList struct {
	imageList        []*cluster.NamespaceImage
	includeNamespace bool
}

func (it IterableImageList) Iter() chan []any {
	c := make(chan []any)
	go func() {
		for _, p := range it.imageList {
			if it.includeNamespace {
				c <- []any{p.Name, p.Namespace, p.Database, p.CreatedAt.AsTime().Local()}
			} else {
				c <- []any{p.Name, p.Database, p.CreatedAt.AsTime().Local()}
			}
		}
		close(c)
	}()
	return c
}

func runImageList(ctx context.Context, cmd *cobra.Command, namespace string) error {
	query := &cluster.GetNamespaceImageListQuery{Namespace: namespace}
	resp, err := client.NewClient().GetCluster().GetNamespaceImageList(context.Background(), query)
	if err != nil {
		cmd.Printf("[error] unable to get list %s\n", err.Error())
		return err
	}

	if len(namespace) == 0 {
		var cols = []any{"NAME", "NAMESPACE", "DATABASE", "CREATED"}
		printConsole(resp, cols, true)
	} else {
		var cols = []any{"NAME", "DATABASE", "CREATED"}
		printConsole(resp, cols, false)
	}

	return nil
}

func printConsole(resp *cluster.GetNamespaceImageListResponse, cols []any, includeNamespace bool) {
	var rows IterableImageList = IterableImageList{
		imageList:        resp.ImageList,
		includeNamespace: includeNamespace,
	}

	var format string
	if includeNamespace {
		format = "%s\t%s\t%s\t%s\n"
	} else {
		format = "%s\t%s\t%s\n"
	}
	console.Print(format, cols, rows)
}
