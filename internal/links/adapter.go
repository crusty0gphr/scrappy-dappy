package links

import (
	"io"
	"log"
	"net/http"
	"sync"

	"scrappy-dappy/internal/domain"
)

type Links interface {
	ExtractValue(body io.Reader, tag, attr string) []string
}

type Adapter struct {
	extractor Links
}

func New(e Links) *Adapter {
	return &Adapter{
		extractor: e,
	}
}

func (a Adapter) Extract(url string, wg *sync.WaitGroup, out chan domain.Output) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	// a.extractor.ExtractValue(resp.Body, "a", "href")
	out <- domain.Output{
		domain.OutputNode{
			Website:    url,
			Route:      "",
			StatusCode: resp.StatusCode,
		},
	}
	wg.Done()
}
