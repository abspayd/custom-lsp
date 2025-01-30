package rpc

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
)

func EncodeMessage(msg any) string {
	content, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("Content-Length: %d\r\n\r\n%s", len(content), content)
}

type BaseMessage struct {
	Method string `json:"method"`
}

func DecodeMessage(msg []byte) (string, []byte, error) {
	headers, content, found := bytes.Cut(msg, []byte{'\r', '\n', '\r', '\n'})
	if !found {
		return "", nil, errors.New("Unable to locate separator")
	}

	contentLength := 0
	header_fields := bytes.Split(headers, []byte{'\r', '\n'})
	for _, header_field := range header_fields {
		contentLengthPrefix := []byte("Content-Length: ")
		if bytes.HasPrefix(header_field, contentLengthPrefix) {
			value, found := bytes.CutPrefix(header_field, contentLengthPrefix)
			if !found {
				return "", nil, errors.New("Unable to locate value for content length header")
			}
			parsedValue, err := strconv.Atoi(string(value))
			if err != nil {
				return "", nil, err
			}
			contentLength = int(parsedValue)
		}
	}

	if contentLength <= 0 {
		return "", nil, errors.New(fmt.Sprintf("Invalid content length: %d", contentLength))
	}

	var baseMessage BaseMessage
	if err := json.Unmarshal(content[:contentLength], &baseMessage); err != nil {
		return "", nil, err
	}

	return baseMessage.Method, content[:contentLength], nil
}

// type SplitFunc func(data []byte, atEOF bool) (advance int, token []byte, err error)

func Split(data []byte, _ bool) (advance int, token []byte, err error) {
	header, content, found := bytes.Cut(data, []byte{'\r', '\n', '\r', '\n'})
	if !found {
		return 0, nil, nil
	}

	contentLength := 0
	header_fields := bytes.Split(header, []byte{'\r', '\n'})
	for _, header_field := range header_fields {
		contentLengthPrefix := []byte("Content-Length: ")
		if bytes.HasPrefix(header_field, contentLengthPrefix) {
			value, _ := bytes.CutPrefix(header_field, contentLengthPrefix)
			parsedValue, err := strconv.Atoi(string(value))
			if err != nil {
				return 0, nil, err
			}
			contentLength = int(parsedValue)
		}
	}

	if len(content) < contentLength {
		return 0, nil, nil
	}

	totalLength := len(header) + 4 + contentLength // header + separator + content
	return totalLength, data[:totalLength], nil
}
