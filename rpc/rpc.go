package rpc

import (
	"encoding/json"
	"fmt"
)

type Message struct {
	JsonRPC string   `json:"jsonrpc"`
	Id      int      `json:"id"`
	Method  string   `json:"method"`
	Params  []string `json:"params,omitempty"`
}

func FormatMessage(msg *Message) {
	if msg.JsonRPC == "" {
		msg.JsonRPC = "2.0"
	}
}

func ValidateMessage(msg Message) error {
	if msg.JsonRPC != "2.0" && msg.JsonRPC != "" {
		return fmt.Errorf(`Invalid jsonrpc: "%s"`, msg.JsonRPC)
	}
	return nil
}

func EncodeMessage(msg Message) string {
	if err := ValidateMessage(msg); err != nil {
		panic(err)
	}
	FormatMessage(&msg)
	content, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("Content-Length: %d\r\n\r\n%s", len(content), content)
}
