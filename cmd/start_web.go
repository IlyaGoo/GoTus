package cmd

import (
	"GoTus/web_server"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(startWeb)
}

var webServer web_server.WebServer

var startWeb = &cobra.Command{
	Use:   "startWeb",
	Short: "Start web server",
	Long:  "Start web server",
	Run: func(cmd *cobra.Command, args []string) {
		webServer = web_server.NewWebServer()
		webServer.StartWebServer()
	},
}
