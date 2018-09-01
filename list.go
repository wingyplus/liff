package main

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/wingyplus/liff/internal/liff"
)

func report(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}

var listcmd = &cobra.Command{
	Use:   "list",
	Short: "Gets information on all the LIFF apps registered in the channel.",
	Run: func(cmd *cobra.Command, args []string) {
		liff.SetAccessToken(viper.GetString("access_token"))
		apps, err := liff.ListApps()
		if err != nil {
			report(err)
		}

		printer := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)

		fmt.Fprintf(printer, "LIFF ID\tVIEW TYPE\tVIEW URL\tLIFF SCHEME\n")
		for _, app := range apps {
			fmt.Fprintf(printer, "%s\t%s\t%s\t%s\n", app.LiffID, app.View.Type, app.View.URL, fmt.Sprintf("line://app/%s", app.LiffID))
		}

		printer.Flush()
	},
}

func init() {
	rootcmd.AddCommand(listcmd)
}
