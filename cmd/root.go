/*
Copyright © 2025 Franek Suszko <franeksu@gmail.com>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var (
	cfgFile string
)

var rootCmd = &cobra.Command{
	Use:   "notifly",
	Short: "Monitor flight price drops",
	Long:  `Notifly helps you track flight prices and get notifications when they drop.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.config/notifly/config.yaml)")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		if err != nil {
			os.Exit(1)
		}
		viper.AddConfigPath(home + "/.config/notifly")
		viper.SetConfigName("config")
	}
	viper.AutomaticEnv()
	viper.ReadInConfig()
}
