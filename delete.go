package main

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/wingyplus/liff/internal/liff"
)

var LiffID string

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
	deletecmd.Flags().StringVar(&LiffID, "liff-id", "", "ID of the LIFF app to be deleted")
	deletecmd.MarkFlagRequired("liff-id")
	rootcmd.AddCommand(deletecmd)
}
