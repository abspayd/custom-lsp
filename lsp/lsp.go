package lsp

import (
	"bufio"
	"custom-lsp/rpc"
	"fmt"
	"os"
	"strings"
)

func Start() {
	// Quick test
	out := bufio.NewWriter(os.Stdout)
	in := bufio.NewReader(os.Stdin)

	// Send initial message
	for {
		str, err := in.ReadString('\n')
		// scanner.Scan()
		if err != nil {
			// Close out the program
			out.Flush()
            return
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
	}
}
