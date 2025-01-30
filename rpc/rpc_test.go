package rpc_test

import (
	"testing"

	"custom-lsp/rpc"
)

type EncodingExample struct {
	Method string
}

func TestEncode(t *testing.T) {
	expected := "Content-Length: 17\r\n\r\n{\"Method\":\"test\"}"
	actual := rpc.EncodeMessage(EncodingExample{Method: "test"})
	if expected != actual {
		t.Fatalf("Expected: %s, Actual: %s", expected, actual)
	}
}

func TestDecode(t *testing.T) {
	incomingMessage := "Content-Length: 17\r\n\r\n{\"Method\":\"test\"}"
	method, content, err := rpc.DecodeMessage([]byte(incomingMessage))
	if err != nil {
		t.Fatal(err)
	}
	if len(content) != 17 {
		t.Fatalf("Expected: 17, Actual: %d", len(content))
	}

	if method != "test" {
		t.Fatalf("Expected: test, Actual: %s", method)
	}
}
