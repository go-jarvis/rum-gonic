package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
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

func (c *Client) Do(ctx context.Context, op operator.APIOperator, meta Meta) (*Result, error) {

	req, err := newRequest(ctx, op, meta)
	if err != nil {
		return nil, err
	}
	req.Header = http.Header(meta)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	return &Result{
		response: resp,
	}, nil
}

func newRequest(ctx context.Context, api operator.APIOperator, meta Meta) (*http.Request, error) {

	if ctx == nil {
		ctx = context.Background()
	}

	b, err := json.Marshal(api)
	if err != nil {
		return nil, err
	}
	buf := bytes.NewBuffer(b)

	req, err := http.NewRequestWithContext(ctx, api.Method(), api.Path(), buf)
	if err != nil {
		return nil, err
	}

	req.Header = http.Header(meta)

	return req, nil
}

type Result struct {
	response *http.Response
}

func (r *Result) Bind(receiver interface{}) error {
	if receiver == nil {
		return nil
	}

	b, err := io.ReadAll(r.response.Body)

	if err != nil {
		return err
	}
	defer r.response.Body.Close()

	fmt.Println(string(b))

	return json.Unmarshal(b, receiver)
}