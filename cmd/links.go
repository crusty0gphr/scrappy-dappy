package cmd

import (
	linksAdapter "scrappy-dappy/internal/links"
	"scrappy-dappy/internal/services/extractor"
	linksClient "scrappy-dappy/pkg/links"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

const (
	flagExtract = "extract-links"
)

var urls []string
var links = &cobra.Command{
	Use:           "links",
	SilenceUsage:  true,
	SilenceErrors: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		// defer func() {
		// 	if r := recover(); r != nil {
		// 		log.Printf("Recovered. Error: %s", r)
		// 	}
		// }()

		client := linksClient.New()
		adapter := linksAdapter.New(client)
		extractor := extractor.New(adapter)

		err := extractor.Run()
		if err != nil {
			return errors.Wrap(err, ErrServiceError)
		}
		return nil
	},
}
