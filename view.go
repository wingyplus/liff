package main

import "github.com/spf13/pflag"

var (
	ViewType string
	ViewURL  string
)

func viewBindFlags(flags *pflag.FlagSet) {
	flags.StringVar(&ViewType, "view-type", "", "Size of the LIFF app view. Specify one of the following values")
	flags.StringVar(&ViewURL, "view-url", "", "URL of the LIFF app. The URL scheme must be https.")
}
