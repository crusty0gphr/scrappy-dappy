package console

import (
	"fmt"
)

type Client struct {
}

func New() *Client {
	return &Client{}
}

func (c Client) Make(data string) error {
	fmt.Println(data) // this is temporary
	return nil
}
