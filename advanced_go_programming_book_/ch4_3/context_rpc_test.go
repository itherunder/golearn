package ch4_3

import (
	"fmt"
	"net"
	"net/rpc"
	"testing"

	"github.com/yezihack/colorlog"
)

type HelloContextService struct {
	conn    net.Conn
	isLogin bool
}

func (p *HelloContextService) Hello(request string, reply *string) error {
	*reply = fmt.Sprintf("hello:%s, from:%s", request, p.conn.RemoteAddr().String())
	return nil
}

func (p *HelloContextService) Login(request string, reply *string) error {
	if request != "user:password" {
		return fmt.Errorf("auth failed")
	}
	colorlog.Info("login ok")
	*reply = "login ok"
	p.isLogin = true
	return nil
}

func TestContextRpc(t *testing.T) {
	colorlog.Info("listening...")
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		colorlog.Error("listen error: %v", err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			colorlog.Error("accept error: %v", err)
		}
		colorlog.Info("accepted: %+v", conn)

		go func() {
			defer conn.Close()

			p := rpc.NewServer()
			p.Register(&HelloContextService{conn: conn})
			p.ServeConn(conn)
		}()
	}
}

func TestContextRpcClient(t *testing.T) {
	client, err := rpc.Dial("tcp", ":1234")
	if err != nil {
		colorlog.Error("dial error: %v", err)
	}

	var reply string
	err = client.Call("HelloContextService.Hello", "ither", &reply)
	if err != nil {
		colorlog.Error("call error: %v", err)
	}

	colorlog.Info("reply: %q", reply)
	client.Close()
}
