package main

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/wingyplus/liff/internal/liff"
)

var updatecmd = &cobra.Command{
	Use:   "update",
	Short: "Adds an app to LIFF. You can add up to 30 LIFF apps on one channel.",
	Run: func(cmd *cobra.Command, args []string) {
		liff.SetAccessToken(viper.GetString("access_token"))
		fmt.Println("finding LIFF application...")
		apps, err := liff.ListApps()
		if err != nil {
			report(err)
		}
		for _, app := range apps {
			if LiffID == app.LiffID {
				fmt.Println("found LIFF application, updating...")
				vt := ViewType
				if vt == "" {
					vt = app.View.Type
				}
				vu := ViewURL
				if vu == "" {
					vu = app.View.URL
				}
				fmt.Println(vt, vu)
				if err := liff.Update(LiffID, &liff.View{Type: vt, URL: vu}); err != nil {
					report(err)
				}
				fmt.Println("done")
				return
			}
		}
		report(errors.New(fmt.Sprintf("Liff id %s not found", LiffID)))
	},
}

func init() {
	bindViewFlags(updatecmd, false, false)
	bindLiffIDFlag(updatecmd)
	rootcmd.AddCommand(updatecmd)
}
