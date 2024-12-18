package rpc_test

import (
	"fmt"
	"testing"
)

func TestDecode(t *testing.T) {
	content := "{\"method\": \"hello\"}"
	header := fmt.Sprintf("Content-Length: %d\r\n\r\n", len(content))
	req := []byte(header + content)

	// fmt.Println(string(req))

	lsp_request, err := DecodeMessage(req)
	if err != nil {
		t.Fatalf("Unable to decode request: %v", err)
	}

	if lsp_request.Method != "hello" {
		t.Fatalf("Mismatched value: %v", err)
	}
}

func TestEncode(t *testing.T) {
    msg_text := "Hello"
	msg := Message{
		Header:  MessageHeader{
            ContentLength: len(fmt.Sprintf("{\"message\":\"Hello\"}")),
        },
		Content: MessageContent{
            Message: "Hello",
        },
	}

    _ = msg_text
    encoded_message := EncodeMessage(msg)
    fmt.Printf("%v\n", encoded_message)
    t.Fatal()
}
