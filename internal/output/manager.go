package output

import (
	"github.com/intel-go/fastjson"
	"scrappy-dappy/internal/domain"
)

const (
	FormatJSON = "json"
)

type Console interface {
	Make(data string) error
}

type File interface {
	Make(data []byte, path, extension string) error
}

type Manager struct {
	console Console
	file    File
}

func New(console Console, file File) *Manager {
	return &Manager{
		console: console,
		file:    file,
	}
}

func (m Manager) Make(data domain.Output, path string, outputType domain.OutputType) (err error) {
	jsonB, err := fastjson.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	switch outputType {
	case domain.OutputTypeConsole:
		err = m.console.Make(string(jsonB))
	case domain.OutputTypeJson:
		err = m.file.Make(jsonB, path, FormatJSON)
	}
	return
}
