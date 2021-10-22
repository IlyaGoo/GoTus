package cmd

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// Used for flags.
	cfgFile string
	port    string

	rootCmd = &cobra.Command{
		Use:   "GoTus",
		Short: "A test project for organizing parties",
		Long:  `A test project for organizing parties`,
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobra.yaml)")
	rootCmd.PersistentFlags().StringVarP(&port, "port", "p", "", "port of web server")
	// Bind viper flags to cobra
	viper.BindPFlag("port", rootCmd.PersistentFlags().Lookup("port"))
	viper.SetDefault("port", "8080")

	rootCmd.AddCommand(startWeb)
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag
		viper.SetConfigFile(cfgFile)
	} else {
		// Get current directory
		path, err := os.Getwd()
		cobra.CheckErr(err)

		// Search config in home directory with name ".config" (without extension)
		viper.AddConfigPath(path)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".config")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}

	initLogger()
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
