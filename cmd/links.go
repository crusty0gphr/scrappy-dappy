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
	flagExtract = "extract"
)

var urls []string
var links = &cobra.Command{
	Use:           "links",
	SilenceUsage:  true,
	SilenceErrors: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("Panic! - Error: %s", r)
			}
		}()

		client := htmlClient.New()
		adapter := linksExtractor.New(client)
		extractor := extractor.New(adapter)

		if err := extractor.Run(urls); err != nil {
			return errors.Wrap(err, ErrService)
		}
		return nil
	},
}
