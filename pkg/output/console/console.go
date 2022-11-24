package console

import (
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
)

var (
	header = table.Row{"#", "website", "route", "status", "error"}
)

type Node struct {
	Website    string
	Route      string
	StatusCode int
	Error      error
}

type Input []Node

type Client struct {
	t table.Writer
}

func New() *Client {
	return &Client{
		t: table.NewWriter(),
	}
}

func (c *Client) Make(in Input) error {
	c.makeTable(in)
	c.t.Render()
	return nil
}

func (c *Client) makeTable(in Input) {
	c.t.SetOutputMirror(os.Stdout)
	c.t.AppendHeader(header)

	for i, d := range in {
		var errMsg string
		if d.Error != nil {
			errMsg = d.Error.(error).Error()
		}

		row := table.Row{i, d.Website, d.Route, d.StatusCode, errMsg}
		c.t.AppendRow(row)
		c.t.AppendSeparator()
	}
}
