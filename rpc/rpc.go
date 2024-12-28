package rpc

import (
	"encoding/json"
	"fmt"
)

// References:
//  https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification
//  https://www.jsonrpc.org/specification

type (
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

func Encode[T Request | Response](msg T) ([]byte, error) {
	encoded_message, err := json.Marshal(msg)
	if err != nil {
		return nil, err
	}
    encoded_message = []byte(fmt.Sprintf("Content-Length: %d\r\n\r\n%v", len(encoded_message), encoded_message))
	return encoded_message, nil
}

func Decode[T Response | Request](data []byte) (*T, error) {
	return nil, nil
}

// Read the header and content sections of an RPC and return their values. Returns and error on failure.
// func ReadRequest(r io.Reader) (header Header, content Request, err error) {
// 	scanner := bufio.NewScanner(r)
//
// 	// read the message headers
// 	for scanner.Scan() {
// 		line := scanner.Text()
// 		if line == "" {
// 			break
// 		}
//
// 		field, value, found := strings.Cut(line, ":")
// 		if !found {
// 			return Header{}, Request{}, fmt.Errorf("Invalid header line: %v", header)
// 		}
//
// 		field = strings.ToLower(strings.TrimSpace(field))
// 		switch field {
// 		case "content-length":
// 			str := strings.TrimSpace(value)
// 			content_length, err := strconv.Atoi(str)
// 			if err != nil {
// 				return Header{}, Request{}, err
// 			}
// 			header.ContentLength = content_length
// 			break
// 		case "content-type":
// 			header.ContentType = strings.TrimSpace(value)
// 			break
// 		default:
// 			return Header{}, Request{}, fmt.Errorf("Unknown header: \"%s\"", field)
// 		}
// 	}
//
// 	// read the rest of the message content
// 	content_body := ""
// 	for scanner.Scan() {
// 		content_body += scanner.Text()
// 	}
// 	if err := scanner.Err(); err != nil {
// 		return Header{}, Request{}, err
// 	}
//
// 	err = json.Unmarshal([]byte(content_body), &content)
// 	if err != nil {
// 		return Header{}, Request{}, err
// 	}
//
// 	// TODO: I might want to validate the JsonRPC value just to make sure that it's actually v2.0.
//
// 	return header, content, nil
// }

// func (r *Response) Send(w io.Writer) error {
// 	encoded_response, err := Encode(*r)
// 	if err != nil {
// 		return err
// 	}
// 	r.JsonRPC = "2.0"
// 	w.Write([]byte(encoded_response))
// 	return nil
// }

// func Encode[M Request | Response](msg M) (string, error) {
// 	content, err := json.Marshal(msg)
// 	if err != nil {
// 		return "", err
// 	}
// 	return fmt.Sprintf("Content-Length: %d\r\n\r\n%s", len(content), content), nil
// }

// func DecodeRequest(data []byte) (Request, error) {
// 	_, content, found := bytes.Cut(data, []byte("\r\n\r\n"))
// 	if !found {
// 		return Request{}, fmt.Errorf("Unable to find message content: %s", data)
// 	}
//
// 	var request Request
// 	err := json.Unmarshal(content, &request)
// 	if err != nil {
// 		return Request{}, err
// 	}
//
// 	return request, nil
// }
