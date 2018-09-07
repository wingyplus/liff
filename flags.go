package main

import (
	"github.com/spf13/cobra"
)

var LiffID string

func bindLiffIDFlag(cmd *cobra.Command) {
	cmd.Flags().StringVar(&LiffID, "liff-id", "", "ID of the LIFF app to be deleted")
	cmd.MarkFlagRequired("liff-id")
}

var (
	ViewType string
	ViewURL  string
)

func bindViewFlags(cmd *cobra.Command, requiredViewType, requiredViewURL bool) {
	cmd.Flags().StringVar(&ViewType, "view-type", "", "Size of the LIFF app view. Specify one of the following values")
	cmd.Flags().StringVar(&ViewURL, "view-url", "", "URL of the LIFF app. The URL scheme must be https.")

	if requiredViewType {
		addcmd.MarkFlagRequired("view-type")
	}
	if requiredViewURL {
		addcmd.MarkFlagRequired("view-url")
	}
}
