package image

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/quarkloop/quarkloop/cli/client"
	v1 "github.com/quarkloop/quarkloop/pkg/grpc/daemon/v1"
	grpc "github.com/quarkloop/quarkloop/pkg/grpc/daemon/v1/image"
)

func NewImageRemoveCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "rm",
		Short: "Remove a image in cluster",
		RunE: func(cmd *cobra.Command, args []string) error {
			namespace, _ := cmd.Flags().GetString("namespace")
			image, _ := cmd.Flags().GetString("name")

			return runImageRemove(cmd.Context(), cmd, namespace, image)
		},
	}

	flags := cmd.Flags()
	flags.String("namespace", "", "image namespace")
	flags.String("name", "", "image name to remove")

	cmd.MarkFlagRequired("namespace")
	cmd.MarkFlagRequired("name")

	return cmd
}

func runImageRemove(ctx context.Context, cmd *cobra.Command, namespace, image string) error {
	deleteCmd := &grpc.DeleteImageCommand{
		Metadata: &v1.Metadata{
			Namespace: namespace,
			Name:      image,
		},
	}
	_, err := client.NewClient().GetImage().DeleteImage(context.Background(), deleteCmd)
	if err != nil {
		cmd.Printf("[error] unable to create %s\n", err.Error())
		return err
	}

	cmd.Printf("Image deleted: %s\n", image)
	return nil
}
