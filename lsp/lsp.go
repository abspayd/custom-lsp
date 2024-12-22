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
	return "Rename!", nil
}

func context() (string, error) {
	return "", nil
}

func Start() {
	for {
		_, content, err := rpc.ReadRequest(os.Stdin)
		if err != nil {
			panic(fmt.Sprintf("Unable to read request: %v", err))
		}

		// TODO: process received request
		id := content.Id // keep this for the response
		method := content.Method

		f, ok := accepted_methods[method]
		if !ok {
			// respond with error method not found (rpc.MethodNotFound)
			err_msg := rpc.ResponseError{
				Code:    rpc.MethodNotFound,
				Message: fmt.Sprintf("Unknown method: \"%s\"", method),
			}
			response := &rpc.Response{
				JsonRPC: "2.0",
				Id:      id,
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
            panic(fmt.Sprintf("Method %s returned with an error: %v", method, err))
        }

		response := &rpc.Response{
			JsonRPC: "2.0",
			Id:      id,
			Result:  res,
		}
		encoded_response, err := rpc.Encode(response)
		if err != nil {
			panic(fmt.Sprintf("Unable to encode response: %v", err))
		}

		fmt.Print(encoded_response)

		// TODO: response
		// response := &rpc.Response{
		//     JsonRPC: "2.0",
		//     Id: id,
		//     Result: ,
		// }

		os.Stdout.Close()
		return // loop one time (for now)
	}

	// TODO: is the method call case insensitive in the LSP spec?

	// TODO: loop two previous steps (list and process)

	/*
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
	*/
}
