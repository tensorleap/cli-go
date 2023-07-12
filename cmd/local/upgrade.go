package local

import (
	"errors"

	"github.com/spf13/cobra"
	"github.com/tensorleap/cli-go/pkg/helm"
	"github.com/tensorleap/cli-go/pkg/k3d"
	"github.com/tensorleap/cli-go/pkg/local"
)

func NewUpgradeCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "upgrade",
		Short: "Upgrade an existing local tensorleap installation to the latest version",
		Long:  `Upgrade an existing local tensorleap installation to the latest version`,
		RunE: func(cmd *cobra.Command, args []string) error {

			if err := local.ValidateVarDir(); err != nil {
				return err
			}
			ctx := cmd.Context()
			bgLogger := local.SetupBackgroundLogger("upgrade")

			cluster, err := local.GetTensorleapCluster(ctx)
			if err != nil {
				return err
			}
			isGpuCluster := k3d.IsGpuCluster(cluster)

			helmConfig, err := helm.CreateHelmConfig(local.KUBE_CONTEXT, local.KUBE_NAMESPACE, bgLogger)
			if err != nil {
				return err
			}

			isHelmReleaseExisted, err := helm.IsHelmReleaseExists(helmConfig)
			if err != nil {
				return err
			} else if !isHelmReleaseExisted {
				return errors.New("Not found helm release, Please make sure to install before upgrade")
			}

			imagesToCache, imageToCacheInTheBackground, err := local.GetLatestImages(isGpuCluster)
			if err != nil {
				return err
			}

			registry, err := k3d.GetRegistry(ctx)
			if err != nil {
				return err
			}
			k3d.CacheImagesInParallel(ctx, imagesToCache, registry)

			if err := helm.UpgradeTensorleapChartVersion(helmConfig); err != nil {
				return err
			}

			k3d.CacheImageInTheBackground(ctx, imageToCacheInTheBackground)
			return nil
		},
	}
	return cmd
}

func init() {
	RootCommand.AddCommand(NewUpgradeCmd())
}
