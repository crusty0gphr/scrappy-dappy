package links

import (
	"io"
	"net/http"
	"sync"

	"scrappy-dappy/internal/domain"
)

type Links interface {
	ExtractValueByAttrName(body io.Reader, tag, attr string) []string
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
	defer wg.Done()

	resp, err := http.Get(url)
	if err != nil {
		out <- domain.Output{
			domain.OutputNode{
				Website:    url,
				Route:      url,
				StatusCode: http.StatusBadRequest,
				Err:        err,
			},
		}
		return
	}

	result := make(domain.Output, 0)
	links := a.extractor.ExtractValueByAttrName(resp.Body, "a", "href")
	result = append(
		result, domain.OutputNode{
			Website:    url,
			Route:      url,
			StatusCode: resp.StatusCode,
		},
	)

	for _, link := range links {
		r, err := http.Get(link)
		result = append(
			result, domain.OutputNode{
				Website:    url,
				Route:      link,
				StatusCode: r.StatusCode,
				Err:        err,
			},
		)
	}
	out <- result
}
