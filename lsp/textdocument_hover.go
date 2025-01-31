package lsp

type HoverRequest struct {
	Request
	Params HoverParams `json:"params"`
}

type HoverParams struct {
	TextDocumentPositionParams
}

type HoverResponse struct {
	Response
	Result HoverResult `json:"result"`
}

type HoverResult struct {
	Contents MarkupContent `json:"contents"`
}

type MarkupContent struct {
	/**
	 * The type of the Markup
	 */
	Kind string `json:"kind"`

	/**
	 * The content itself
	 */
	Value string `json:"value"`
}

func NewHoverResponse(id int, contents string) HoverResponse {
	return HoverResponse{
		Response: Response{
			RPC: "2.0",
			ID:  &id,
		},
		Result: HoverResult{
			Contents: MarkupContent{
				Kind:  "plaintext",
				Value: contents,
			},
		},
	}
}
