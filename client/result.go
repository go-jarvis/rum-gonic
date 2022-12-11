package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Result struct {
	response *http.Response
}

func (r *Result) Bind(receiver interface{}) (Meta, error) {
	meta := Meta(r.response.Header)
	if receiver == nil {
		return meta, nil
	}

	b, err := io.ReadAll(r.response.Body)

	if err != nil {
		return meta, err
	}
	defer r.response.Body.Close()

	fmt.Println(string(b))

	return meta, json.Unmarshal(b, receiver)
}
