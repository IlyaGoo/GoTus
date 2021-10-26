package cmd

import (
	"GoTus/web_server"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(startWeb)
}

var webPresenter web_server.WebServer

var startWeb = &cobra.Command{
	Use:   "startWeb",
	Short: "Start web server",
	Long:  "Start web server",
	Run: func(cmd *cobra.Command, args []string) {
		webPresenter = web_server.NewWebServer()
		webPresenter.StartWebServer()
	},
}
