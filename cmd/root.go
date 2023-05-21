package cmd

import (
	"fmt"
	"os"
  "path"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tensorleap/cli-go/cmd/auth"
	"github.com/tensorleap/cli-go/cmd/datasets"
	"github.com/tensorleap/cli-go/cmd/local"
	"github.com/tensorleap/cli-go/cmd/models"
)

var RootCommand = &cobra.Command{
	Use:   "tensorleap",
	Short: "Tensorleap - Deepbug your models!",
	Long: `A debugger and analyzer for your DNNs.
Complete documentation is available at http://docs.tensoleap.ai`,
}

var cfgFile string

func init() {
  cobra.OnInitialize(initConfig)
	RootCommand.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.config/tensorleap/config.yaml)")
	RootCommand.AddCommand(auth.RootCommand)
	RootCommand.AddCommand(local.RootCommand)
  RootCommand.AddCommand(datasets.RootCommand)
  RootCommand.AddCommand(models.RootCommand)
}

func Execute() {
	if err := RootCommand.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		homeDir, err := os.UserHomeDir()
		cobra.CheckErr(err)

    configDir := path.Join(homeDir, ".config", "tensorleap")
		viper.AddConfigPath(configDir)
		viper.SetConfigType("yaml")
		viper.SetConfigName("config")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}

