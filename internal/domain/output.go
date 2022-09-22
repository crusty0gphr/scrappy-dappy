package domain

type Output []OutputNode

type OutputNode struct {
	Website    string `json:"website"`
	Route      string `json:"route"`
	StatusCode int    `json:"statusCode"`
	Err        error  `json:"error,omitempty"`
}
