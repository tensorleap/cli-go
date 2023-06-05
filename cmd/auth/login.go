package auth

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	. "github.com/tensorleap/cli-go/pkg/api"
	"github.com/tensorleap/cli-go/pkg/config"
)

var cmd = &cobra.Command{
	Use:   "login [api key] [api url]",
	Args:  cobra.ExactArgs(2),
	Short: "Loging with Tensorleap API key",
	Long:  `Loging with Tensorleap API key`,
	RunE: func(cmd *cobra.Command, args []string) error {
		apiKey := args[0]
		baseUrl := args[1]
		ctx := CreateAuthenticatedContext(cmd.Context(), baseUrl, apiKey)

		userData, _, err := ApiClient.WhoAmI(ctx).Execute()
		if err != nil {
			return err
		}
		fmt.Println("User email: " + userData.Local.Email)
		fmt.Println("Team name: " + userData.OrganizationName)

		viper.Set(config.API_URL_CONFIG_PATH, baseUrl)
		viper.Set(config.API_KEY_CONFIG_PATH, apiKey)
		err = viper.SafeWriteConfig()
		if err != nil {
			_, ok := err.(viper.ConfigFileAlreadyExistsError)
			if ok {
				err = viper.WriteConfig()
			}
		}
		if err != nil {
			return err
		}
		fmt.Println("Saved credentials to: ", viper.ConfigFileUsed())
		return nil
	},
}

func init() {
	RootCommand.AddCommand(cmd)
}
