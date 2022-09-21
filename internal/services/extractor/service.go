package extractor

import (
	"fmt"
	"sync"

	"scrappy-dappy/internal/domain"
)

type LinksAdapter interface {
	Extract(w string, wg *sync.WaitGroup, out chan domain.Output)
}

type Service struct {
	adapter LinksAdapter
}

func New(a LinksAdapter) *Service {
	return &Service{
		adapter: a,
	}
}

func (s Service) Run(websites []string) error {
	var wg sync.WaitGroup
	out := make(chan domain.Output)

	for _, website := range websites {
		wg.Add(1)
		go s.adapter.Extract(website, &wg, out)
	}

	result := make([]domain.Output, 0)
	for i := 0; i < len(websites); i++ {
		v := <-out
		result = append(result, v)
	}
	wg.Wait()

	fmt.Println(result)
	return nil
}
