package ch4_3

import (
	"net"
	"net/rpc"
	"testing"
	"time"

	"github.com/yezihack/colorlog"
)

type HelloService struct{}

func (p *HelloService) Hello(request string, reply *string) error {
	*reply = "hello:" + request
	return nil
}

func TestReverseRpc(t *testing.T) {
	rpc.Register(new(HelloService))

	for {
		conn, _ := net.Dial("tcp", ":1234")
		if conn == nil {
			time.Sleep(time.Second)
			continue
		}
		// go func() {
		// 	rpc.ServeConn(conn)
		// 	conn.Close()
		// }()
		rpc.ServeConn(conn)
		conn.Close()
	}
}

func TestReverseRpcClient(t *testing.T) {
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		colorlog.Error("listen error: %v", err)
	}

	clientChan := make(chan *rpc.Client)

	go func() {
		for {
			colorlog.Info("accepting...")
			conn, err := listener.Accept()
			if err != nil {
				colorlog.Error("accept error: %v", err)
			}
			colorlog.Info("conn: %+v", conn)

			clientChan <- rpc.NewClient(conn)
		}
	}()

	doClientWork(clientChan)
}

func doClientWork(clientChan <-chan *rpc.Client) {
	client := <-clientChan
	defer client.Close()

	var reply string
	err := client.Call("HelloService.Hello", "ither", &reply)
	if err != nil {
		colorlog.Error("call error: %v", err)
	}

	colorlog.Info("reply: %q", reply)
}
