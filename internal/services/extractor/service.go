package extractor

import (
	"errors"
)

type LinksAdapter interface {
	Extract() error
}

type Service struct {
	adapter LinksAdapter
}

func New(a LinksAdapter) *Service {
	return &Service{
		adapter: a,
	}
}

func (s Service) Run() error {
	// panic("fuck off! not implemented")
	return errors.New("err, fuck")
}
