package rpc_test

import (
	"custom-lsp/rpc"
	"encoding/json"
	"fmt"
	"io"
	"reflect"
	"testing"
)

func TestEncodeResponse(t *testing.T) {
	msg := rpc.Response{
		JsonRPC: "2.0",
		Id:      1,
		Result:  "textDocument/completion",
	}

	expected_content := `{"jsonrpc":"2.0","id":1,"result":"textDocument/completion"}`
	expected_message := fmt.Sprintf("Content-Length: %d\r\n\r\n%s", len(expected_content), expected_content)

	result_message, err := rpc.Encode(msg)
	if err != nil {
		t.Fatal(err)
	}
	if result_message != expected_message {
		t.Fatalf("%#v != %#v", result_message, expected_message)
	}
}

func TestEncodeRequest(t *testing.T) {
	msg := rpc.Request{
		JsonRPC: "2.0",
		Id:      1,
		Method:  "textDocument/completion",
	}

	expected_content := `{"jsonrpc":"2.0","id":1,"method":"textDocument/completion"}`
	expected_message := fmt.Sprintf("Content-Length: %d\r\n\r\n%s", len(expected_content), expected_content)

	result_message, err := rpc.Encode(msg)
	if err != nil {
		t.Fatal(err)
	}
	if result_message != expected_message {
		t.Fatalf("%#v != %#v", result_message, expected_message)
	}
}

func TestDecodeRequest(t *testing.T) {
	request_content := &rpc.Request{
		JsonRPC: "2.0",
		Method:  "exit",
	}

	request_message, err := rpc.Encode(*request_content)
	if err != nil {
		t.Fatal(err)
	}
	decoded_request, err := rpc.DecodeRequest([]byte(request_message))
	if err != nil {
		t.Fatal(err)
	}
	original_request := *request_content

	if original_request != decoded_request {
		t.Fatalf("%#v != %#v", original_request, decoded_request)
	}

	if !reflect.DeepEqual(original_request, decoded_request) {
		t.Fatalf("%#v != %#v", original_request, decoded_request)
	}
}

func TestReadHeader(t *testing.T) {
	request := rpc.Request{
		JsonRPC: "2.0",
		Method:  "exit",
	}
	content, err := json.Marshal(request)
	if err != nil {
		t.Fatal(err)
	}
	encoded_request, err := rpc.Encode(request)

	r, w := io.Pipe()

	// write to the pipe
	go func() {
		defer w.Close()
		w.Write([]byte(encoded_request))
	}()

	headers, _, err := rpc.ReadRequest(r)
	if err != nil {
		t.Fatal(err)
	}

	if headers.ContentLength != len(content) {
		t.Fatalf("Header got content length = %d, expected %d.", headers.ContentLength, len(content))
	}
}

func TestReadContent(t *testing.T) {
	request := rpc.Request{
		JsonRPC: "2.0",
		Method:  "exit",
	}
	encoded_request, err := rpc.Encode(request)

	r, w := io.Pipe()

	go func() {
		defer w.Close()
		w.Write([]byte(encoded_request))
	}()

	// read the headers to move the reader forward (does this work?)
	_, c, err := rpc.ReadRequest(r)
	if err != nil {
		t.Fatal(err)
	}
    
    if c.JsonRPC != request.JsonRPC || c.Id != request.Id || c.Method != request.Method {
        t.Fatalf("Got %#v, expected %#v", c, request)
    }
}
