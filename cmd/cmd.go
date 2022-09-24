package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{}

func init() {
	rootCmd.AddCommand(links)
	links.PersistentFlags().StringSliceVar(
		&urls,
		flagExtract,
		nil,
		"extract links form the given websites.\nusage: --extract=\"https://example.com,https://example2.com\"",
	)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
