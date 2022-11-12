package console

import (
	"os"

	"github.com/intel-go/fastjson"
	"github.com/jedib0t/go-pretty/v6/table"
)

var (
	header = table.Row{"#", "website", "route", "status", "error"}
)

type data struct {
	Website    string
	Route      string
	StatusCode int
	Error      error
}

type input []data

type Client struct {
	t table.Writer
}

func New() *Client {
	return &Client{
		t: table.NewWriter(),
	}
}

func (c *Client) Make(data string) error {
	var r input
	err := fastjson.Unmarshal([]byte(data), &r)
	if err != nil {
		return err
	}

	c.makeTable(r)
	c.t.Render()
	return nil
}

func (c *Client) makeTable(in input) {
	c.t.SetOutputMirror(os.Stdout)
	c.t.AppendHeader(header)

	for i, d := range in {
		row := table.Row{i, d.Website, d.Route, d.StatusCode}
		c.t.AppendRow(row)
		c.t.AppendSeparator()
	}
}
