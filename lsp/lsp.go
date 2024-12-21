package lsp

import (
	"bufio"
	"custom-lsp/rpc"
	"fmt"
	"io"
	"os"
	"strings"
)

func readHeader(r io.Reader) ([]string, error) {
	reader := bufio.NewReader(r)

    header := []string{} // TODO: make this a map instead?
	var last_byte byte
	for reader.Buffered() > 0 {
		b, err := reader.ReadByte()
		if err != nil {
			return nil, err
		}
		if b == '\r' {
			// end of header section; add header to header string slice
			continue
		} else if b == '\n' {
			if last_byte == '\n' {
				return header, nil
			}
			continue
		}

		last_byte = b
	}

	return nil, nil
}

func readContent() string {
	return ""
}

func Start() {

	readHeader(os.Stdin)

	// Quick test
	out := bufio.NewWriter(os.Stdout)
	in := bufio.NewReader(os.Stdin)

	// Send initial message
	for {
		str, err := in.ReadString('\n')
		if err != nil {
			// unexpected error
			os.Exit(1)
		}

		str, found := strings.CutSuffix(str, "\n")
		if !found {
			panic("This should be impossible")
		}

		msg := &rpc.Response{
			Id:     1,
			Result: str,
		}
		rsp, err := rpc.Encode(msg)

		fmt.Fprint(out, rsp)
		out.Flush()

		// just loop once (for now)
		os.Stdout.Close()
		return
	}
}
