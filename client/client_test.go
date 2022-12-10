package client

import (
	"context"
	"fmt"
	"testing"

	"github.com/go-jarvis/rum-gonic/pkg/httpx"
	"github.com/go-jarvis/rum-gonic/pkg/operator"
)

func TestNewClient(t *testing.T) {

	c := NewClient()

	hp := &httpbin{}

	data := &Data{}
	re, err := c.Do(context.TODO(), hp)
	if err != nil {
		panic(err)
	}

	err = re.Bind(data)
	if err != nil {
		panic(err)
	}
	fmt.Println("data=", data)

}

var _ operator.APIOperator = &httpbin{}

type httpbin struct {
	httpx.MethodGet
}

func (hp *httpbin) Path() string {
	return "https://httpbin.org/get"
}

type Data struct {
	Args    map[string]any    `json:"args"`
	Headers map[string]string `json:"headers"`
	Origin  string            `json:"origin"`
	URL     string            `json:"url"`
}
