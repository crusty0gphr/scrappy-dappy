package file

import (
	"fmt"
	"os"

	"github.com/google/uuid"
	"github.com/intel-go/fastjson"
	"github.com/pkg/errors"
)

const (
	perms = 0644
)

var (
	ErrInvalidPath      = errors.New("invalid path")
	ErrInvalidExtension = errors.New("invalid extension")
)

var (
	extensionJson = "json"
)

var supportedExtensions = map[string]struct{}{
	extensionJson: {},
}

type Node struct {
	Error      error
	Website    string
	Route      string
	StatusCode int
}

type Input []Node

type Client struct {
}

func New() *Client {
	return &Client{}
}

func (c Client) Make(in Input, path, extension string) error {
	b, err := fastjson.MarshalIndent(in, "", "  ")
	if err != nil {
		return err
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return ErrInvalidPath
	}
	if !supportedExtension(extension) {
		return ErrInvalidExtension
	}

	if last := len(path) - 1; last >= 0 && path[last] == '/' {
		path = path[:last]
	}

	name := uuid.New()
	dir := fmt.Sprintf("%s/%s.%s", path, name, extension)

	err = os.WriteFile(dir, b, perms)
	if err != nil {
		return err
	}
	return nil
}

func supportedExtension(extension string) bool {
	if _, ok := supportedExtensions[extension]; !ok {
		return false
	}
	return true
}
