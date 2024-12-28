package lsp

import (
	"custom-lsp/rpc"
	"fmt"
	"os"
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


var (
	accepted_methods = map[string]func() (string, error){
		"textDocument/rename":  Rename,
		"textDocument/context": Context,
	}
)

type Server struct {
    Initialized bool
}

func Rename() (string, error) {
	return "TODO: Rename", nil
}

func Context() (string, error) {
	return "TODO: Context", nil
}

func (server Server) Initialize() (string, error) {
    if server.Initialized {
        return "", fmt.Errorf("Server is already initialized")
    }

    // TODO: listen for initialize request
    _, content, err := rpc.ReadRequest(os.Stdin)
    if err != nil {
        return "", err
    }
    if content.Method != "initialize" {
        // TODO: return error
    }

    // TODO

	return "", nil
}

func (server Server) Error(message rpc.ResponseError) error {

    return nil
}

func (server Server) Listen() {
    if !server.Initialized {
        // TODO: Send error message and exit.
    }
}

func (server Server) Exit() {

}

// Depricated: this is getting replaced by Server.Listen()
func Start() {
    initialized := false
	for {
		_, content, err := rpc.ReadRequest(os.Stdin)
		if err != nil {
			panic(fmt.Sprintf("Unable to read request: %v", err))
		}

        if !initialized && content.Method != "initialize" {

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
