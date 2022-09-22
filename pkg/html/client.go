package html

import (
	"io"
	"net/url"

	"golang.org/x/net/html"
)

type Client struct {
}

func New() *Client {
	return &Client{}
}

func (c Client) ExtractValueByAttrName(body io.Reader, tag, attr string) (output []string) {
	draft := make(map[string]struct{})
	t := html.NewTokenizer(body)

tokenParser:
	for {
		tokenType := t.Next()
		switch tokenType {
		case html.ErrorToken:
			break tokenParser // jump out of this mess!
		case html.StartTagToken, html.EndTagToken:
			token := t.Token()
			if token.Data != tag {
				continue
			}
			for _, attribute := range token.Attr {
				if attribute.Key != attr {
					continue
				}
				// avoid duplicates
				if _, ok := draft[attribute.Val]; !ok && isUrl(attribute.Val) {
					draft[attribute.Val] = struct{}{} // saving little space here, for no reason... idk
				}
			}
		}
	}
	for key := range draft {
		output = append(output, key)
	}
	return
}

func isUrl(s string) bool {
	u, err := url.Parse(s)
	return err == nil && u.Scheme != "" && u.Host != ""
}
