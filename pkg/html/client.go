package html

import (
	"io"
	
	"golang.org/x/net/html"
)

type Client struct {
}

func New() *Client {
	return &Client{}
}

func (c Client) ExtractValue(body io.Reader, tag, attr string) (r []string) {
	// draft := make(map[string]string)
	t := html.NewTokenizer(body)
	
	for {
		token := t.Next()
		switch token {
		
		}
	}
}
