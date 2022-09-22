package extractor

import (
	"fmt"
	"github.com/intel-go/fastjson"
	"log"
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

func (s Service) Run(websites []string, depth uint) error {
	var layer uint
	var wg sync.WaitGroup
	defer wg.Wait()

	log.Print("service started")
	out := make(chan domain.Output)
	defer close(out)

	result := make(domain.Output, 0)
	for layer <= depth {
		for _, website := range websites {
			wg.Add(1)
			go s.links.Extract(website, &wg, out)
		}
		for i := 0; i < len(websites); i++ {
			result = append(result, <-out...)
		}
		websites = s.outputToInput(result)
		layer++
	}

	b, err := fastjson.MarshalIndent(result, "", "  ")
	if err != nil {
		return err
	}

	fmt.Println(string(b))
	log.Print("service finished")
	return nil
}

func (s Service) outputToInput(o domain.Output) (res []string) {
	for _, outputNode := range o {
		res = append(res, outputNode.Route)
	}
	return
}
