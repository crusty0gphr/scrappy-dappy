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
	links.PersistentFlags().UintVar(
		&depth,
		flagDepthLimiter,
		25,
		"limit looping depth of extracted links.\nusage: --depth=10 (default 25)\n 0 means scan only homepage",
	)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
