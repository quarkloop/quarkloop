package image

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/quarkloop/quarkloop/cli/client"
	"github.com/quarkloop/quarkloop/cli/console"
	v1 "github.com/quarkloop/quarkloop/pkg/grpc/daemon/v1/image"
)

func NewImageListCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ls",
		Short: "List namespace images",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runImageList(cmd.Context(), cmd)
		},
	}

	return cmd
}

type IterableImageList []*v1.Image

func (it IterableImageList) Iter() chan []any {
	c := make(chan []any)
	go func() {
		for _, p := range it {
			c <- []any{p.Name, p.CreatedAt.AsTime().Local()}
		}
		close(c)
	}()
	return c
}

func runImageList(ctx context.Context, cmd *cobra.Command) error {
	resp, err := client.NewClient().GetImage().GetImageList(context.Background(), nil)
	if err != nil {
		cmd.Printf("[error] unable to get list %s\n", err.Error())
		return err
	}

	var cols = []any{"NAME", "CREATED"}
	var rows IterableImageList = resp.Images
	console.Print("%s\t%s\t%s\n", cols, rows)

	return nil
}
