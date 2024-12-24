package lsp

import (
	"custom-lsp/rpc"
	"fmt"
	"os"
)

var (
	accepted_methods = map[string]func() (string, error){
		"textDocument/rename":  rename,
		"textDocument/context": context,
	}
)

func rename() (string, error) {
	return "TODO: Rename", nil
}

func context() (string, error) {
	return "TODO: Context", nil
}

func Start() {
	for {
		_, content, err := rpc.ReadRequest(os.Stdin)
		if err != nil {
			panic(fmt.Sprintf("Unable to read request: %v", err))
		}

		switch content.Method {
		case "shutdown":
			// this is sent as a request
			break
		case "exit":
			// this is sent as a notification
            os.Stdout.Close()
            return
		}

		f, ok := accepted_methods[content.Method]
		if !ok {
			// respond with error method not found (rpc.MethodNotFound)
			err_msg := rpc.ResponseError{
				Code:    rpc.MethodNotFound,
				Message: fmt.Sprintf("Unknown method: \"%s\"", content.Method),
			}
			response := rpc.Response{
				JsonRPC: "2.0",
				Id:      content.Id,
				Error:   &err_msg,
			}

			// write the response
			encoded_response, err := rpc.Encode(response)
			if err != nil {
				panic(fmt.Sprintf("Unable to encode response: %v", err))
			}

			fmt.Print(encoded_response)
		}

		// execute the method
		res, err := f()
		if err != nil {
			panic(fmt.Sprintf("Method %s returned with an error: %v", content.Method, err))
		}

		response := rpc.Response{
			JsonRPC: "2.0",
			Id:      content.Id,
			Result:  res,
		}
		encoded_response, err := rpc.Encode(response)
		if err != nil {
			panic(fmt.Sprintf("Unable to encode response: %v", err))
		}

		fmt.Print(encoded_response)
	}
}
