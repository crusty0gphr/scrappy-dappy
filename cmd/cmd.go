package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var cmd = &cobra.Command{}

func init() {
	cmd.AddCommand(links)
	links.PersistentFlags().StringSliceVar(
		&urls,
		flagExtract,
		nil,
		"extract links form the given websites.\nusage: --extract-links=\"example.com,example2.com\"",
	)
}

func Execute() {
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
