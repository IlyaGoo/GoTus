package cmd

import (
	"GoTus/web_presenter"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	rootCmd.AddCommand(startWeb)
}

var startWeb = &cobra.Command{
	Use:   "startWeb",
	Short: "Start web server",
	Long:  "Start web server",
	Run: func(cmd *cobra.Command, args []string) {
		testWebPresenter := web_presenter.WebPresenter{viper.GetString("port"), ""}
		testWebPresenter.StartWebPresenter()
	},
}
