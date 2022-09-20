package links

type Links interface {
	Extract(websites []string)
}

type Adapter struct {
	extractor Links
}

func New(e Links) *Adapter {
	return &Adapter{
		extractor: e,
	}
}

func (a Adapter) Extract() error {
	panic("fuck off! not implemented")
}
