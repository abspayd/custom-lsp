package lsp

type DidChangeTextDocumentNotification struct {
	Notification
	Params DidChangeTextDocumentParams `json:"params"`
}

type DidChangeTextDocumentParams struct {
	TextDocument   VersionedTextDocumentIdentifier `json:"textDocument"`
	ContentChanges []TextDocumentChangeEvent       `json:"contentChanges"`
}

type TextDocumentChangeEvent struct {
	Text string `json:"text"`
}
