package main

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/wingyplus/liff/internal/liff"
)

var addcmd = &cobra.Command{
	Use:   "add",
	Short: "Adds an app to LIFF. You can add up to 30 LIFF apps on one channel.",
	Run: func(cmd *cobra.Command, args []string) {
		liff.SetAccessToken(viper.GetString("access_token"))
		fmt.Println("adding app to LIFF...")
		id, err := liff.Add(&liff.View{
			Type: ViewType,
			URL:  ViewURL,
		})
		if err != nil {
			report(err)
		}
		fmt.Printf("LIFF id: %s", id)
		fmt.Println("done")
	},
}

func init() {
	bindViewFlags(addcmd, true, true)
	rootcmd.AddCommand(addcmd)
}
