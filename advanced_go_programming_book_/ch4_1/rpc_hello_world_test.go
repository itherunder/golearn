package ch4_1

import (
	"net"
	"net/rpc"
	"testing"
	"time"

	"github.com/yezihack/colorlog"
)

type HelloService struct{}

// 其中Hello方法必须满足Go语言的RPC规则：
// 方法只能有两个可序列化的参数，其中第二个参数是指针类型，
// 并且返回一个error类型，同时必须是公开的方法。
func (p *HelloService) Hello(request string, reply *string) error {
	*reply = "hello:" + request
	return nil
}

func TestRpcHelloWorld(t *testing.T) {
	rpc.RegisterName("HelloService", new(HelloService))

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		colorlog.Error("ListenTcp error: %v", err)
	}
	colorlog.Info("Listening...")

	conn, err := listener.Accept()
	if err != nil {
		colorlog.Error("Accept error: %v", err)
	}
	defer conn.Close()

	colorlog.Info("Accepted: %+v", conn)

	rpc.ServeConn(conn)
	time.Sleep(20 * time.Second)
}

func TestClientRpcRequest(t *testing.T) {
	client, err := rpc.Dial("tcp", ":1234")
	if err != nil {
		colorlog.Error("Dial rpc server error: %v", err)
	}

	var reply string
	err = client.Call("HelloService.Hello", "ither", &reply)
	if err != nil {
		colorlog.Error("Call rpc method error: %v", err)
	}

	colorlog.Info("reply: %q", reply)
}
