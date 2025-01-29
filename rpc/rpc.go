package rpc

type Request struct {
	JsonRPC string        `json:"jsonrpc"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
	Id      interface{}   `json:"id"`
}

type Response struct {
	JsonRPC string      `json:"jsonrpc"`
	Id      interface{} `json:"id"`
	Result  string      `json:"result,omitempty"`
	Error   *Error      `json:"error,omitempty"`
}

type Error struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

/*
type Server struct {
	Reader bufio.Reader
	Writer bufio.Writer
}

func NewServer(r io.Reader, w io.Writer) Server {
	return Server{
		Reader: *bufio.NewReader(r),
		Writer: *bufio.NewWriter(w),
	}
}

func (server Server) Close() error {
	return server.Writer.Flush()
}

func (server Server) Start() {
	defer server.Close()

	for {
		line, err := server.Reader.ReadBytes(io.EOF)
	}
}
*/
