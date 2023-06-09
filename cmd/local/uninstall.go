package local

import (
	"github.com/spf13/cobra"
	"github.com/tensorleap/cli-go/pkg/k3d"
	"github.com/tensorleap/cli-go/pkg/local"
	"github.com/tensorleap/cli-go/pkg/log"
)

func NewUninstallCmd() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "uninstall",
		Short: "Remove local Tensorleap installation",
		Long:  `Remove local Tensorleap installation`,
		RunE: func(cmd *cobra.Command, args []string) error {
			log.SetCommandName("uninstall")
			log.SendCloudReport("info", "Starting uninstall", "Starting", &map[string]interface{}{"args": args})
			close, err := local.SetupInfra("uninstall")
			if err != nil {
				return err
			}
			defer close()

			ctx := cmd.Context()
			err = k3d.UninstallCluster(ctx)
			if err != nil {
				return err
			}

			log.SendCloudReport("info", "Successfully completed uninstall", "Success", nil)
			return nil
			// todo: implement --cache and --data
		},
	}

	return cmd
}

func init() {
	RootCommand.AddCommand(NewUninstallCmd())
}
