package rpc

// References:
//  https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification
//  https://www.jsonrpc.org/specification

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"strconv"
	"strings"
)

const (
	// Defined by JSON-RPC
	ParseError     = -32700
	InvalidRequest = -32600
	MethodNotFound = -32601
	InvalidParams  = -32602
	InternalError  = -32603

	// Start range of JSON-RPC reserved error codes
	JsonRpcReservedErrorRangeStart = -32099

	// Error code indicating that a server received a notification or
	// request before the server has received the `initialize` request.
	ServerNotInitialized = -32002
	UnknownErrorCode     = -32001

	// End of JSON-RPC reserved error code range
	JsonRpcReservedErrorRangeEnd = -32000

	RequestFailed            = -32803
	ServerCancelled          = -32802
	ContentModified          = -32801
	RequestCancelled         = -32800
	LspReservedErrorRangeEnd = -32800
)

type (
	Message interface {
        Send(io.Writer) error
	}

	// RPC message header
	Header struct {
		ContentLength int
		ContentType   string
	}

	Request struct {
		JsonRPC string `json:"jsonrpc"`
		Id      int    `json:"id"`
		Method  string `json:"method"`
		Params  any    `json:"params,omitempty"`
	}

	Response struct {
		JsonRPC string         `json:"jsonrpc"`
		Id      int            `json:"id,omitempty"`
		Result  any            `json:"result,omitempty"` // MUST be empty on error
		Error   *ResponseError `json:"error,omitempty"`  // Error object in case a request fails
	}

	ResponseError struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Data    any    `json:"data,omitempty"`
	}
)

// Read the header and content sections of an RPC and return their values. Returns and error on failure.
func ReadRequest(r io.Reader) (header Header, content Request, err error) {
	scanner := bufio.NewScanner(r)

	// read the message headers
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		field, value, found := strings.Cut(line, ":")
		if !found {
			return Header{}, Request{}, fmt.Errorf("Invalid header line: %v", header)
		}

		field = strings.ToLower(strings.TrimSpace(field))
		switch field {
		case "content-length":
			str := strings.TrimSpace(value)
			content_length, err := strconv.Atoi(str)
			if err != nil {
				return Header{}, Request{}, err
			}
			header.ContentLength = content_length
			break
		case "content-type":
			header.ContentType = strings.TrimSpace(value)
			break
		default:
			return Header{}, Request{}, fmt.Errorf("Unknown header: \"%s\"", field)
		}
	}

	// read the rest of the message content
	content_body := ""
	for scanner.Scan() {
		content_body += scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		return Header{}, Request{}, err
	}

	err = json.Unmarshal([]byte(content_body), &content)
	if err != nil {
		return Header{}, Request{}, err
	}

	// TODO: I might want to validate the JsonRPC value just to make sure that it's actually v2.0.

	return header, content, nil
}

func (r *Response) Send(w io.Writer) error {
	encoded_response, err := Encode(*r)
	if err != nil {
		return err
	}
	r.JsonRPC = "2.0"
	w.Write([]byte(encoded_response))
	return nil
}

func (r *Response) Format() {
	r.JsonRPC = "2.0"
}

func (r *Request) Format() {
	r.JsonRPC = "2.0"
}

func Encode[M Request | Response](msg M) (string, error) {
	content, err := json.Marshal(msg)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("Content-Length: %d\r\n\r\n%s", len(content), content), nil
}

func DecodeRequest(data []byte) (Request, error) {
	_, content, found := bytes.Cut(data, []byte("\r\n\r\n"))
	if !found {
		return Request{}, fmt.Errorf("Unable to find message content: %s", data)
	}

	var request Request
	err := json.Unmarshal(content, &request)
	if err != nil {
		return Request{}, err
	}

	return request, nil
}
