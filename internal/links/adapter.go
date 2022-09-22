package links

import (
	"io"
	"log"
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

	log.Printf("started extracting %s", url)
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

	var statusCode int
	for _, link := range links {
		// shadowed variables, same name but inside the different scope, so nothing to worry about :)
		resp, err := http.Get(link)
		if err != nil {
			statusCode = http.StatusBadRequest
		} else {
			statusCode = resp.StatusCode
		}

		result = append(
			result, domain.OutputNode{
				Website:    url,
				Route:      link,
				StatusCode: statusCode,
				Err:        err,
			},
		)
	}
	out <- result
	log.Printf("finished extracting %s", url)
}
