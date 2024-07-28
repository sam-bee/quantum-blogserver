package cmd

import (
	"fmt"
	"os"
	"sam-bee/quantum-blogserver/pkg/quantumblogserver"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var confPath string
var config quantumblogserver.ApplicationConfig

var rootCmd = &cobra.Command{
	Use:   "quantumblogerver",
	Short: "Run webserver",
	Long:  `Run webserver to host a static site using post-quantum cryptography`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVarP(&confPath, "config", "c", "", "config file (required)")
}

func initConfig() {
	if _, err := os.Stat(confPath); os.IsNotExist(err) {
		os.Stderr.WriteString("Config file not found\n")
		os.Exit(1)
	}

	viper.SetConfigType("env")
	viper.SetConfigFile(confPath)

	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}

	viper.AutomaticEnv()

	config = quantumblogserver.BuildConfig(viper.AllSettings())
}
