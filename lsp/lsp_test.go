package lsp

import (
	"fmt"
	"testing"
)

func TestDecode(t *testing.T) {
	content := "{\"method\": \"hello\"}"
	header := fmt.Sprintf("Content-Length: %d\r\n\r\n", len(content))
    req := []byte(header + content)

    fmt.Println(string(req))

    lsp_request, err := Decode(req)
    if err != nil {
        t.Fatalf("Unable to decode request: %v", err)
    }

    if lsp_request.Method != "hello" {
        t.Fatalf("Mismatched value: %v", err)
    }
}
