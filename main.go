package main

import (
	"bufio"
	"custom-lsp/rpc"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Quick test
	out := bufio.NewWriter(os.Stdout)
	in := bufio.NewReader(os.Stdin)

	// Send initial message
    fmt.Println("LSP Listening on stdin...")
	out.Flush()

	for {
		str, err := in.ReadString('\n')
		if err != nil {
            // Close out the program
            out.Flush()
            os.Exit(0)
		}
        str, found := strings.CutSuffix(str,"\n")
        if !found {
            panic("This should be impossible")
        }


		msg := rpc.Message{
			Id:     1,
			Method: str,
			Params: nil,
		}

		rpc_msg := rpc.EncodeMessage(msg)

		fmt.Printf("Message received: %s\n", str)
		fmt.Fprint(out, rpc_msg)
		out.Flush()
	}
}
