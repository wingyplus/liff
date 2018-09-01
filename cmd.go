package main

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootcmd = &cobra.Command{
	Use:   "liff [ACTION]",
	Short: "Command-Line utility for LIFF application.",
}

func init() {
	rootcmd.PersistentFlags().String("access-token", "", "Access token for channel")
	viper.SetEnvPrefix("LIFF")
	viper.AutomaticEnv()
	viper.BindPFlag("access_token", rootcmd.PersistentFlags().Lookup("access-token"))
}

// Execute command-line
func Execute() error {
	return rootcmd.Execute()
}
