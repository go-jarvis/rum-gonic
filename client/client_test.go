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
	meta := Meta{
		"TraceId": []string{"trace_id", "abc"},
		"Span-ID": []string{"spanID", "span123"},
	}

	re, err := c.Do(context.TODO(), hp, meta)
	if err != nil {
		panic(err)
	}

	data := &Data{}
	_, err = re.Bind(data)
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
