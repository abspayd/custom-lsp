package rpc_test

import (
	"custom-lsp/rpc"
	"fmt"
	"reflect"
	"testing"
)

func TestEncodeResponse(t *testing.T) {
	msg := &rpc.Response{
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
	msg := &rpc.Request{
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

func TestEncodeNotification(t *testing.T) {
	msg := &rpc.Notification{
		JsonRPC: "2.0",
		Method:  "textDocument/completion",
	}

	expected_content := `{"jsonrpc":"2.0","method":"textDocument/completion"}`
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
		Id:      1,
		Method:  "textDocument/completion",
		Params:  nil,
	}

	request_message, err := rpc.Encode(request_content)
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
