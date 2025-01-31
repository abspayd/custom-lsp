package main

import (
	"bufio"
	"custom-lsp/analysis"
	"custom-lsp/lsp"
	"custom-lsp/rpc"
	"encoding/json"
	"log"
	"os"
)

func main() {
	homedir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	logger := getLogger(homedir + "/Developer/custom-lsp/log.txt")
	logger.Println("LSP started")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(rpc.Split)

	state := analysis.NewState()

	for scanner.Scan() {
		msg := scanner.Bytes()
		method, contents, err := rpc.DecodeMessage(msg)
		if err != nil {
			logger.Printf("Got an error: %s", err)
			continue
		}

		handleMessage(logger, state, method, contents)
	}
}

func handleMessage(logger *log.Logger, state analysis.State, method string, contents []byte) {
	logger.Printf("Received msg with method: %s", method)

	switch method {
	case "initialize":
		var request lsp.InitializeRequest
		if err := json.Unmarshal(contents, &request); err != nil {
			logger.Printf("initialize: %s", err)
			return
		}

		logger.Printf("Connected to: %s %s", request.Params.ClientInfo.Name, *request.Params.ClientInfo.Version)

		// send response
		msg := lsp.NewInitializeResponse(request.ID)
		reply := rpc.EncodeMessage(msg)

		writer := os.Stdout
		writer.Write([]byte(reply))

		logger.Print("Sent a reply")
	case "textDocument/didOpen":
		var notification lsp.DidOpenTextDocumentNotification
		if err := json.Unmarshal(contents, &notification); err != nil {
			logger.Printf("textDocument/didOpen: %s", err)
			return
		}

		logger.Printf("Opened: %s", notification.Params.TextDocument.URI)
		state.OpenDocument(notification.Params.TextDocument.URI, notification.Params.TextDocument.Text)
	case "textDocument/didChange":
		var notification lsp.DidChangeTextDocumentNotification
		if err := json.Unmarshal(contents, &notification); err != nil {
			logger.Printf("textDocument/didChange: %s", err)
			return
		}

		logger.Printf("Changed: %s", notification.Params.TextDocument.URI)
		for _, change := range notification.Params.ContentChanges {
			state.UpdateDocument(notification.Params.TextDocument.URI, change.Text)
		}
	case "textDocument/hover":
		var request lsp.HoverRequest
		if err := json.Unmarshal(contents, &request); err != nil {
			logger.Printf("textDocument/hover: %s", err)
			return
		}

		logger.Printf("Hover: %s {%d, %d}", request.Params.TextDocument, request.Params.Position.Character, request.Params.Position.Line)

		// send response
		msg := lsp.NewHoverResponse(request.ID, "Hello")
		reply := rpc.EncodeMessage(msg)

		writer := os.Stdout
		writer.Write([]byte(reply))

		logger.Print("Sent a reply")
	}
}

func getLogger(filename string) *log.Logger {
	logfile, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	return log.New(logfile, "[custom-lsp] ", log.Ldate|log.Ltime|log.Lshortfile)
}
