package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"strings"

	"github.com/go-jarvis/rum-gonic/pkg/operator"
	"github.com/go-jarvis/rum-gonic/pkg/reflectx"
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

func newRequest(ctx context.Context, op operator.APIOperator, meta Meta) (*http.Request, error) {

	if ctx == nil {
		ctx = context.Background()
	}

	b, err := json.Marshal(op)
	if err != nil {
		return nil, err
	}
	buf := bytes.NewBuffer(b)

	path := requestPath(op)
	fmt.Println(path)

	req, err := http.NewRequestWithContext(ctx, op.Method(), path, buf)
	if err != nil {
		return nil, err
	}

	req.Header = http.Header(meta)

	return req, nil
}

func requestPath(op operator.APIOperator) string {
	path := op.Path()
	queries := []string{}

	rv := reflect.ValueOf(op)
	rv = reflectx.Indirect(rv)

	rt := reflect.TypeOf(op)
	rt = reflectx.Deref(rt)

	for i := 0; i < rt.NumField(); i++ {
		fv := rv.Field(i)
		ft := rt.Field(i)
		path = replacePath(path, ft, fv)
		query := parseQuery(ft, fv)
		if len(query) != 0 {
			queries = append(queries, query)
		}
	}

	if len(queries) != 0 {
		path = fmt.Sprintf("%s?%s", path, strings.Join(queries, "&"))
	}

	return path
}

func replacePath(path string, ft reflect.StructField, fv reflect.Value) string {

	tag := ft.Tag.Get("uri")
	if tag == "" {
		return path
	}

	fv = reflectx.Indirect(fv)
	value := ""
	switch v := fv.Interface().(type) {
	case string, *string:
		value = fmt.Sprint(v)
	case int, int8, int16, int32, int64,
		uint, uint8, uint16, uint32, uint64:
		value = fmt.Sprint(v)
	case *int, *int8, *int16, *int32, *int64,
		*uint, *uint8, *uint16, *uint32, *uint64:
		value = fmt.Sprint(v)
	case bool, *bool:
		value = fmt.Sprint(v)
	}

	// path=/user/:name
	path = strings.ReplaceAll(path, ":"+tag, value)
	// path=/user/*name
	path = strings.ReplaceAll(path, "*"+tag, value)

	return path
}

func parseQuery(ft reflect.StructField, fv reflect.Value) string {
	name, ok := ft.Tag.Lookup("query")
	if !ok {
		return ""
	}

	// value := ""
	switch val := fv.Interface().(type) {
	case string:
		return fmt.Sprintf("%s=%s", name, val)
	case *string:
		return fmt.Sprintf("%s=%s", name, *val)
	case int, int8, int16, int32, int64,
		uint, uint8, uint16, uint32, uint64:
		return fmt.Sprintf("%s=%d", name, val)
	case *int:
		return fmt.Sprintf("%s=%d", name, *val)
	case bool:
		v := strconv.FormatBool(val)
		return fmt.Sprintf("%s=%s", name, v)
	case *bool:
		v := strconv.FormatBool(*val)
		return fmt.Sprintf("%s=%s", name, v)
	case []string:
		for i, v := range val {
			v = fmt.Sprintf("%s=%s", name, v)
			val[i] = v
		}
		return strings.Join(val, "&")
	case []*string:
		valc := make([]string, len(val))
		for i, v := range val {
			valc[i] = fmt.Sprintf("%s=%s", name, *v)
		}
		return strings.Join(valc, "&")
	case []int:
		valc := make([]string, len(val))
		for i, v := range val {
			valc[i] = fmt.Sprintf("%s=%d", name, v)
		}
		return strings.Join(valc, "&")
	case []*int:
		valc := make([]string, len(val))
		for i, v := range val {
			valc[i] = fmt.Sprintf("%s=%d", name, *v)
		}
		return strings.Join(valc, "&")
	}

	return ""
}
