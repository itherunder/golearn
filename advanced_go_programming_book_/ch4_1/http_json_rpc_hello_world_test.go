package ch4_1

import (
	"io"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
	"testing"

	"github.com/yezihack/colorlog"
)

func TestHttpJsonRpcHelloWorld(t *testing.T) {
	err := RegistrerHelloService(new(HelloService))
	if err != nil {
		colorlog.Error("Register hello service error: %v", err)
	}
	http.HandleFunc("/jsonrpc", func(w http.ResponseWriter, r *http.Request) {
		var conn io.ReadWriteCloser = struct {
			io.Writer
			io.ReadCloser
		}{
			w,
			r.Body,
		}
		rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	})

	http.ListenAndServe(":1234", nil)
}
