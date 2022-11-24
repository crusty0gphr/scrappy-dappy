package cmd

import (
	"log"

	linksManager "scrappy-dappy/internal/links"
	outputManager "scrappy-dappy/internal/output"
	"scrappy-dappy/internal/services/extractor"
	htmlClient "scrappy-dappy/pkg/html"
	consoleClient "scrappy-dappy/pkg/output/console"
	fileClient "scrappy-dappy/pkg/output/file"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

const (
	flagExtract = "extract"
	flagOutput  = "output"
	flagPath    = "path"
)

var urls []string
var outputType string
var path string
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

		linksManager := linksManager.New(
			htmlClient.New(),
		)
		outputManager := outputManager.New(
			consoleClient.New(),
			fileClient.New(),
		)
		extractor := extractor.New(linksManager, outputManager)

		if err := extractor.Run(urls, outputType, path); err != nil {
			return errors.Wrap(err, ErrService)
		}
		return nil
	},
}
