package rpc

// References:
//  https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification
//  https://www.jsonrpc.org/specification

import (
	"bytes"
	"encoding/json"
	"fmt"
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
		Format()
	}

	Request struct {
		JsonRPC string `json:"jsonrpc"` // MUST be "2.0"
		Id      int    `json:"id"`
		Method  string `json:"method"`
		Params  any    `json:"params,omitempty"`
	}

	Notification struct {
		JsonRPC string   `json:"jsonrpc"` // MUST be "2.0"
		Method  string   `json:"method"`
		Params  []string `json:"params,omitempty"`
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

func (r *Response) Format() {
	r.JsonRPC = "2.0"
}
func (r *Request) Format() {
	r.JsonRPC = "2.0"
}
func (n *Notification) Format() {
	n.JsonRPC = "2.0"
}

func Encode(msg Message) (string, error) {
	msg.Format()
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

func DecodeNotification(data []byte) (Notification, error) {
	_, content, found := bytes.Cut(data, []byte("\r\n\r\n"))
	if !found {
		return Notification{}, fmt.Errorf("Unable to find message content: %s", data)
	}

	var notification Notification
	err := json.Unmarshal(content, &notification)
	if err != nil {
		return Notification{}, err
	}

	return notification, nil
}
