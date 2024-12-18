package rpc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
)

type Request struct {
	// Header part
	ContentLength int `json:"Content-Length"`
	// ContentType   string `json:"Content-Type"`

	// Content part
	JsonRPC string   `json:"jsonrpc"`
	Id      int      `json:"id"`
	Method  string   `json:"method"`
	Params  []string `json:"params"`
}

type Message struct {
	Header  MessageHeader
	Content MessageContent
}

type MessageHeader struct {
	ContentLength int    `json:"Content-Length"`
	ContentType   string `json:"Content-Type,omitempty"`
}
type MessageContent struct {
	JsonRPC string `json:"jsonrpc,omitempty"`
	Id      int    `json:"id,omitempty"`
	Message string `json:"message"`
}

func EncodeMessage(msg Message) string {

	content, err := json.Marshal(msg.Content)
	if err != nil {
		panic(err)
	}

	header := fmt.Sprintf("Content-Length: %d", msg.Header.ContentLength)
	if msg.Header.ContentType != "" {
		header += fmt.Sprintf("\r\nContent-Type: \"%s\"", msg.Header.ContentType)
	}
	header += "\r\n\r\n"

	return header + string(content)
}

func DecodeMessage(data []byte) (Request, error) {
	request := new(Request)

	// Separate header from content body
	headers, content, found := bytes.Cut(data, []byte("\r\n\r\n"))
	if !found {
		panic(fmt.Errorf("Invalid request: %s", string(data)))
	}

	header_fields := bytes.Split(headers, []byte("\r\n"))

	for _, field := range header_fields {
		if !bytes.HasPrefix(field, []byte("Content-Length")) {
			continue // no support for any other headers (atm)
		}

		_, length_field, found := bytes.Cut(field, []byte(":"))
		if !found {
			panic(fmt.Errorf("Malformed content header: %s", string(length_field)))
		}
		length_field = bytes.TrimSpace(length_field)

		content_length, err := strconv.Atoi(string(length_field))
		if err != nil {
			panic(fmt.Errorf("Unable to parse int from Content-Length header: %v", err))
		}
		request.ContentLength = content_length
	}

	if err := json.Unmarshal(content, &request); err != nil {
		panic(fmt.Errorf("Something went wrong when parsing data: %v", err))
	}

	return *request, nil
}