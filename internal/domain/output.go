package domain

type OutputType string

const (
	OutputTypeConsole OutputType = "console"
	OutputTypeJson    OutputType = "file"
)

var OutputTypes = map[OutputType]struct{}{
	OutputTypeConsole: {},
	OutputTypeJson:    {},
}

type Output []OutputNode

type OutputNode struct {
	Err        error  `json:"error,omitempty"`
	Website    string `json:"website"`
	Route      string `json:"route"`
	StatusCode int    `json:"statusCode"`
}
