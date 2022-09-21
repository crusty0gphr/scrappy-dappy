package extractor

import (
	"fmt"
	"github.com/intel-go/fastjson"
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

	out := make(chan domain.Output)

	for _, website := range websites {
		wg.Add(1)
		go s.links.Extract(website, &wg, out)
	}

	result := make(domain.Output, 0)
	for i := 0; i < len(websites); i++ {
		result = append(result, <-out...)
	}

	b, err := fastjson.MarshalIndent(result, "", "  ")
	if err != nil {
		return err
	}

	fmt.Println(string(b))
	return nil
}
