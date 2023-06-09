package local

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/tensorleap/cli-go/pkg/k3d"
	"github.com/tensorleap/cli-go/pkg/log"
)

func ValidateStandaloneDir() error {
	_, err := os.Stat(STANDALONE_DIR)
	if os.IsNotExist(err) {
		log.SendCloudReport("error", "Installation dir not found", "Failed", &map[string]interface{}{"error": err.Error()})
		return errors.New(
			fmt.Sprintf("Not found data directory(%s) on this machine, Please make sure to install before upgrade", STANDALONE_DIR),
		)
	}
	return err
}

func GetTensorleapCluster(ctx context.Context) (*k3d.Cluster, error) {
	cluster, err := k3d.GetCluster(ctx)
	if err != nil {
		log.SendCloudReport("error", "Failed getting k3d cluster", "Failed", &map[string]interface{}{"error": err.Error()})
		return nil, err
	}
	if cluster == nil {
		log.SendCloudReport("error", "K3d cluster found was null", "Failed", &map[string]interface{}{"error": err.Error()})
		return nil, errors.New(
			fmt.Sprintf("Not found local Cluster(%s) on this machine, Please make sure to install before upgrade", STANDALONE_DIR),
		)
	}
	return cluster, nil
}
