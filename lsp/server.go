package lsp

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"custom-lsp/rpc"
)

type Server struct {
	Reader bufio.Reader
	Writer bufio.Writer
}

func NewServer(r io.Reader, w io.Writer) Server {
	return Server{
		Reader: *bufio.NewReader(r),
		Writer: *bufio.NewWriter(w),
	}
}

func (server Server) Start() {
	fmt.Println("LSP server started")
	for {
		// headers
		contentLength := 0
		for {
			line, err := server.Reader.ReadBytes('\n')
			if err != nil {
				panic(fmt.Sprintf("Unable to read bytes: %v", err))
			}

			line = bytes.TrimSpace(line)
			if len(line) == 0 {
				// end of headers
				break
			}

			if bytes.HasPrefix(line, []byte("Content-Length:")) {
				_, value, found := bytes.Cut(line, []byte(": "))
				if !found {
					panic(fmt.Sprintf("Unable to find header field separator for header line: %s", line))
				}
				value_int, err := strconv.ParseInt(string(value), 10, 0)
				if err != nil {
					panic(fmt.Sprintf("Unable to parse Content-Length header: %v", err))
				}

				contentLength = int(value_int)
			}

			// skip other headers
		}

		// content
		if contentLength <= 0 {
			panic("No content provided with header")
		}

		content := make([]byte, contentLength)
		_, err := io.ReadFull(&server.Reader, content)
		if err != nil {
			panic(fmt.Sprintf("Unable to read contents: %v", err))
		}

		rpcRequest := rpc.Request{}
		err = json.Unmarshal(content, rpcRequest)
		if err != nil {
			panic(fmt.Sprintf("Unable to unmarshal content: %v", err))
		}

		rpcResponse, err := handleRequest(rpcRequest)
		if err != nil {
			panic(fmt.Sprintf("Error handling request: %v", err))
		}

		if rpcResponse != nil {
			resp, err := json.Marshal(rpcResponse)
			if err != nil {
				panic(fmt.Sprintf("Unable to marshal server response: %v", err))
			}
			server.Writer.Write(resp)
		}
	}
}

func handleRequest(request rpc.Request) (*rpc.Response, error) {
	return nil, nil
}
