package cmd

import (
	"log"
	"sam-bee/quantum-blogserver/pkg/quantumblogserver"

	"github.com/spf13/cobra"
	"gopkg.in/natefinch/lumberjack.v2"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run webserver",
	Long:  `Run webserver to host a static site using post-quantum cryptography`,
	Run: func(cmd *cobra.Command, args []string) {

		log.SetOutput(&lumberjack.Logger{
			Filename:   config.GetLogFile(),
			MaxSize:    100, // megabytes
			MaxBackups: 3,
			MaxAge:     28,   //days
			Compress:   true, // disabled by default
		})

		quantumblogserver.RunWebserver(config)
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
