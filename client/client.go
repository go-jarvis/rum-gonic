package client

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/go-jarvis/rum-gonic/pkg/operator"
)

type Client struct {
	client *http.Client
}

func NewClient() *Client {
	return &Client{
		client: http.DefaultClient,
	}
}

func (c *Client) Do(ctx context.Context, op operator.APIOperator) (*Result, error) {

	req, err := http.NewRequestWithContext(ctx, op.Method(), op.Path(), nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	return &Result{
		response: resp,
	}, nil
}

type Result struct {
	response *http.Response
}

func (r *Result) Bind(data interface{}) error {
	// if data == nil {
	// 	return nil
	// }

	b, err := ioutil.ReadAll(r.response.Body)
	if err != nil {
		return err
	}
	defer r.response.Body.Close()

	fmt.Println(string(b))

	return json.Unmarshal(b, data)
}
