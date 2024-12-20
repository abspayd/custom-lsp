package lsp_test

import (
	"custom-lsp/rpc"
	"testing"
)

func TestStart(t *testing.T) {
	// client request
	request := rpc.Request{
		Id:     1,
		Method: "textDocument/rename",
		Params: struct {
			Name string `json:"name"`
		}{
			Name: "foo",
		},
	}
	encoded_request, err := rpc.Encode(&request)
	if err != nil {
		t.Fatal(err)
	}

	_ = encoded_request
}
