package links

import (
	"io"
	"log"
	"net/http"
	"net/url"
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

func (a Adapter) Extract(root string, wg *sync.WaitGroup, out chan domain.Output) {
	defer wg.Done()
	result := make(domain.Output, 0)

	log.Printf("started extracting %s", root)
	node, body := a.ping(root)
	node.Website = root
	result = append(result, node)
	if node.Err != nil {
		out <- result
		return
	}
	links := a.extractor.ExtractValueByAttrName(body, "a", "href")
	for _, link := range links {
		if !a.isValidUrl(link) {
			continue
		}
		node, _ := a.ping(link)
		node.Website = root
		result = append(result, node)
	}
	out <- result
	log.Printf("finished extracting %s", root)
}

func (a Adapter) ping(url string) (domain.OutputNode, io.ReadCloser) {
	var statusCode int
	var body io.ReadCloser

	resp, err := http.Get(url)
	if err != nil {
		statusCode = http.StatusBadRequest
	} else {
		statusCode = resp.StatusCode
		body = resp.Body
	}

	return domain.OutputNode{
		Route:      url,
		StatusCode: statusCode,
		Err:        err,
	}, body
}

func (a Adapter) isValidUrl(s string) bool {
	u, err := url.Parse(s)
	return err == nil && u.Scheme != "" && u.Host != ""
}
