package output

import (
	"scrappy-dappy/internal/domain"
	"scrappy-dappy/pkg/output/console"
	"scrappy-dappy/pkg/output/file"
)

const (
	FormatJSON = "json"
)

type Console interface {
	Make(in console.Input) error
}

type File interface {
	Make(in file.Input, path, extension string) error
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

func (m Manager) Make(in domain.Output, path string, outputType domain.OutputType) (err error) {
	switch outputType {
	case domain.OutputTypeConsole:
		err = m.console.Make(
			toConsoleInput(in),
		)
	case domain.OutputTypeJson:
		err = m.file.Make(
			toFileInput(in), path,
			FormatJSON,
		)
	}
	return
}

func toConsoleInput(in domain.Output) (out console.Input) {
	for _, node := range in {
		out = append(out, console.Node{
			Website:    node.Website,
			Route:      node.Route,
			StatusCode: node.StatusCode,
			Error:      node.Err,
		})
	}
	return
}

func toFileInput(in domain.Output) (out file.Input) {
	for _, node := range in {
		out = append(out, file.Node{
			Website:    node.Website,
			Route:      node.Route,
			StatusCode: node.StatusCode,
			Error:      node.Err,
		})
	}
	return
}
