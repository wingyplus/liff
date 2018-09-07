package main

import "github.com/spf13/cobra"

var LiffID string

func bindLiffIDFlag(cmd *cobra.Command) {
	cmd.Flags().StringVar(&LiffID, "liff-id", "", "ID of the LIFF app to be deleted")
	cmd.MarkFlagRequired("liff-id")
}
