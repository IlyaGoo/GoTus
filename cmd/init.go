package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
	"github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// Used for flags.
	configPath string
	port       string

	rootCmd = &cobra.Command{
		Use:   "GoTus",
		Short: "A test project for organizing parties",
		Long:  `A test project for organizing parties`,
	}
)

// Execute executes the root command.
func execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVarP(&configPath, "config", "c", "", "config file (default is ./config/.config)")
	rootCmd.PersistentFlags().StringVarP(&port, "port", "p", "", "port for web-server")
	// Bind viper keys to cobra flags
	viper.BindPFlag("port", rootCmd.PersistentFlags().Lookup("port"))
	viper.SetDefault("port", "8080")
}

func initConfig() {
	if configPath != "" {
		// Use config file from the flag
		viper.SetConfigFile(configPath)
	} else {
		// Get current directory
		path, err := os.Getwd()
		cobra.CheckErr(err)

		// Search config in home directory with name ".config" (without extension)
		viper.AddConfigPath(filepath.Dir(path) + "/configs")
		viper.SetConfigType("yaml")
		viper.SetConfigName(".config")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	} else if _, ok := err.(viper.ConfigFileNotFoundError); ok {
		fmt.Println("Config file not found")
	} else {
		fmt.Println("Config file was found but another error was produced")
	}

	initLogger()

	//Watching re-reading config files
	viper.OnConfigChange(func(e fsnotify.Event) {
		logrus.Info("Config file changed:", e.Name)
	})
	viper.WatchConfig()
}

func initLogger() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout)
	lvl, err := logrus.ParseLevel(viper.GetString("log_level"))
	if err != nil {
		logrus.SetLevel(logrus.DebugLevel)
		logrus.Warn("Using debug level logger")
	} else {
		logrus.SetLevel(lvl)
	}
}
