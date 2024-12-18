package rpc_test

import (
	"custom-lsp/rpc"
	"fmt"
	"testing"
)

func TestEncode(t *testing.T) {
	msg := rpc.Message{
		JsonRPC: "2.0",
		Id:      1,
		Method:  "textDocument/completion",
		Params:  []string{},
	}

	expected_content := `{"jsonrpc":"2.0","id":1,"method":"textDocument/completion"}`
	expected_message := fmt.Sprintf("Content-Length: %d\r\n\r\n%s", len(expected_content), expected_content)

	result_message := rpc.EncodeMessage(msg)
	if result_message != expected_message {
		t.Fatalf("%s != %s", result_message, expected_message)
	}
}
