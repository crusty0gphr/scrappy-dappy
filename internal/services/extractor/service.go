package extractor

import (
	"fmt"
	"github.com/intel-go/fastjson"
	"log"
	"net/http"
	"strings"
	"sync"

	"scrappy-dappy/internal/domain"
)

type LinksAdapter interface {
	Extract(w string, wg *sync.WaitGroup, out chan domain.Output)
}

type Service struct {
	links LinksAdapter
}

func New(a LinksAdapter) *Service {
	return &Service{
		links: a,
	}
}

func (s Service) Run(websites []string) error {
	var wg sync.WaitGroup
	defer wg.Wait()

	log.Print("service started")
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

	jsonB, err := fastjson.MarshalIndent(result, "", "  ")
	if err != nil {
		return err
	}

	fmt.Println(string(jsonB))
	log.Print("service finished")
	return nil
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
