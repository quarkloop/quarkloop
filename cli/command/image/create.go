package image

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/quarkloop/quarkloop/cli/client"
	v1 "github.com/quarkloop/quarkloop/pkg/grpc/v1"
	grpc "github.com/quarkloop/quarkloop/pkg/grpc/v1/image"
)

func NewImageCreateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a image in cluster",
		RunE: func(cmd *cobra.Command, args []string) error {
			namespace, _ := cmd.Flags().GetString("namespace")
			image, _ := cmd.Flags().GetString("name")

			return runImageCreate(cmd.Context(), cmd, namespace, image)
		},
	}

	flags := cmd.Flags()
	flags.String("namespace", "", "image namespace")
	flags.String("name", "", "image name")

	cmd.MarkFlagRequired("namespace")
	cmd.MarkFlagRequired("name")

	return cmd
}

func runImageCreate(ctx context.Context, cmd *cobra.Command, namespace, image string) error {
	createCmd := &grpc.CreateImageCommand{
		Metadata: &v1.Metadata{
			Namespace: namespace,
			Name:      image,
		},
		CreatedBy: "cli",
	}
	_, err := client.NewClient().GetImage().CreateImage(context.Background(), createCmd)
	if err != nil {
		cmd.Printf("[error] unable to create %s\n", err.Error())
		return err
	}

	cmd.Printf("Image created: %s\n", image)
	return nil
}
