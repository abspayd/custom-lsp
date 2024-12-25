package lsp_test

import (
	"bytes"
	"custom-lsp/lsp"
	"custom-lsp/rpc"
	"os"
	"testing"
)

func TestStart(t *testing.T) {
	// client request
	request := rpc.Request{
		Id:     1,
		Method: "exit",
		Params: struct {
			Name string `json:"name"`
		}{
			Name: "foo",
		},
	}
	encoded_request, err := rpc.Encode(request)
	if err != nil {
		t.Fatal(err)
	}

	stdin := os.Stdin
	stdout := os.Stdout
	defer func() {
		os.Stdin = stdin
		os.Stdout = stdout
	}()

	p1r, p1w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	p2r, p2w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}

	os.Stdin = p1r
	go func() {
		// write request
		defer p1w.Close()
		p1w.Write([]byte(encoded_request))
	}()

	os.Stdout = p2w

	lsp.Start()

    var b bytes.Buffer
    n, err := b.ReadFrom(p2r)
    output := make([]byte, n)
    _, err = b.Read(output)
    if err != nil {
        t.Fatal(err)
    }

    // TODO: make this actually test something (e.g. are we getting a valid/expected response back?)
    t.Fatal(string(output))
}

func TestRename(t *testing.T) {

}
