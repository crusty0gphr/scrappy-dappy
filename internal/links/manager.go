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

type Manager struct {
	extractor Links
}

func New(e Links) *Manager {
	return &Manager{
		extractor: e,
	}
}

func (a Manager) Extract(root string, wg *sync.WaitGroup, out chan domain.Output) {
	defer wg.Done()
	result := make(domain.Output, 0)

	node, body := a.ping(root)
	node.Website = root
	result = append(result, node)
	if node.Err != nil {
		out <- result
		return
	}
	a.log("extracting", root)
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
}

func (a Manager) ping(url string) (domain.OutputNode, io.ReadCloser) {
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

func (a Manager) isValidUrl(s string) bool {
	u, err := url.Parse(s)
	return err == nil && u.Scheme != "" && u.Host != ""
}

func (a Manager) log(prefix, s string) {
	const defaultLen = 65
	const separator = "...."

	var msg string
	if len(s) > defaultLen {
		msg = s[:40-3]
		msg += separator
		msg += s[len(s)-40:]
	} else {
		msg = s
	}
	log.Println(prefix, msg)
}
