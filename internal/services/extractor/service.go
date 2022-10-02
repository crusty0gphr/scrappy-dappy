package extractor

import (
	"net/http"
	"strings"
	"sync"

	"github.com/pkg/errors"
	"scrappy-dappy/internal/domain"
)

var (
	ErrInvalidOutputType = errors.New("invalid output type")
)

type LinksManager interface {
	Extract(w string, wg *sync.WaitGroup, out chan domain.Output)
}

type OutputManager interface {
	Make(data domain.Output, path string, outputType domain.OutputType) (err error)
}

type Service struct {
	links  LinksManager
	output OutputManager
}

func New(l LinksManager, o OutputManager) *Service {
	return &Service{
		links:  l,
		output: o,
	}
}

func (s Service) Run(websites []string, outputType, path string) error {
	if _, ok := domain.OutputTypes[domain.OutputType(outputType)]; !ok {
		return ErrInvalidOutputType
	}

	var wg sync.WaitGroup
	defer wg.Wait()

	out := make(chan domain.Output)
	result := make(domain.Output, 0)
	defer close(out)

	for websites != nil {
		for _, website := range websites {
			wg.Add(1)
			go s.links.Extract(strings.TrimSpace(website), &wg, out)
		}
		for i := 0; i < len(websites); i++ {
			result = append(result, <-out...)
		}
		websites = s.outputToInput(websites, result)
	}

	return s.output.Make(result, path, domain.OutputType(outputType))
}

func (s Service) outputToInput(w []string, o domain.Output) (res []string) {
	for _, url := range w {
		for _, outputNode := range o {
			if outputNode.StatusCode != http.StatusOK ||
				url == outputNode.Route ||
				!strings.Contains(outputNode.Route, url) {
				continue
			}
			res = append(res, outputNode.Route)
		}
	}
	return
}
