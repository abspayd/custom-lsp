package lsp

import (
	"custom-lsp/rpc"
	"fmt"
	"os"
)

func Start() {
	headers, content, err := rpc.ReadRequest(os.Stdin)
	if err != nil {
		panic(fmt.Sprintf("Unable to read request: %v", err))
	}

    _ = headers
    _ = content

    


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
