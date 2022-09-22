package cmd

import (
	"log"
	linksExtractor "scrappy-dappy/internal/links"
	"scrappy-dappy/internal/services/extractor"
	htmlClient "scrappy-dappy/pkg/html"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

const (
	flagExtract      = "extract"
	flagDepthLimiter = "depth"
)

var urls []string
var depth uint
var links = &cobra.Command{
	Use:           "links",
	SilenceUsage:  true,
	SilenceErrors: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("Recovered. Error: %s", r)
			}
		}()

		client := htmlClient.New()
		adapter := linksExtractor.New(client)
		extractor := extractor.New(adapter)

		if err := extractor.Run(urls, depth); err != nil {
			return errors.Wrap(err, ErrService)
		}
		return nil
	},
}
