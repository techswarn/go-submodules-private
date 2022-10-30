package main

import (
	// "context"
	// "os"
	// "strings"
	"fmt"
	"github.com/techswarn/tempconv"
)

type Request struct {
	Tempc int `json:tempc`
}

type Response struct {
	StatusCode int               `json:"statusCode,omitempty"`
	Headers    map[string]string `json:"headers,omitempty"`
	Body       string            `json:"body,omitempty"`
}

func Main(in Request) (*Response, error) {
	fmt.Println(in)
	if in.Tempc == 0 {
		fmt.Println("Input is empty")
	}

	c := tempconv.Celsius(in.Tempc)

	return &Response{
		Body: fmt.Sprintf("%v", tempconv.CTof(c)),
	}, nil
}
