package cmd

import (
	"GoTus/web_server"
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(start)
}

var webServer web_server.WebServer

var start = &cobra.Command{
	Use:   "start",
	Short: "Start app",
	Long:  "Start http server etc",
	Run: func(cmd *cobra.Command, args []string) {
		webServer = web_server.NewWebServer()

		go func() {
			if err := webServer.Run(); err != nil {
				logrus.Warn("error occured while running http server: %s", err.Error())
			} else {
				logrus.Info("Http server started")
			}
		}()

		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
		<-quit

		logrus.Info("App Shutting Down")

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer func() {
			cancel()
		}()

		if err := webServer.Shutdown(ctx); err != nil {
			logrus.Fatalf("Http server shutdown failed:%+v", err)
		}

		logrus.Info("Server Exited Properly")
	},
}
