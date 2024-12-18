package main

import (
	"bufio"
	"custom-lsp/rpc"
	"fmt"
	"os"
)

func main() {
	// Quick test
	msg := rpc.Message{
		Id:     1,
		Method: "textDocument/completion",
		Params: nil,
	}

	rpc_msg := rpc.EncodeMessage(msg)
	out := bufio.NewWriter(os.Stdout)
	in := bufio.NewReader(os.Stdin)

	// Send initial message
	fmt.Fprint(out, rpc_msg)

	out.Flush()

	for {
		str, err := in.ReadString('\n')
		if err != nil {
			panic("Unable to read input")
		}

		fmt.Printf("Message received: %s\n", str)
        fmt.Fprint(out, rpc_msg)
        out.Flush()
	}

	// for {
	// 	b := []byte{}
	// 	n, err := in.Read(b)

	// 	if err != nil {
	// 		panic(err)
	// 	}
	//     if n == 0 {
	//         continue
	//     }

	// 	fmt.Printf("Message received: %s\n", string(b))

	// 	fmt.Fprint(out, rpc_msg)
	// 	out.Flush()
	// }
}
