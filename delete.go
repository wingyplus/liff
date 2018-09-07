package main

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/wingyplus/liff/internal/liff"
)

var deletecmd = &cobra.Command{
	Use:   "delete",
	Short: "Deletes a LIFF app.",
	Run: func(cmd *cobra.Command, args []string) {
		liff.SetAccessToken(viper.GetString("access_token"))
		err := liff.Delete(LiffID)
		if err != nil {
			report(err)
		}
		fmt.Println("done")
	},
}

func init() {
	bindLiffIDFlag(deletecmd)
	rootcmd.AddCommand(deletecmd)
}
