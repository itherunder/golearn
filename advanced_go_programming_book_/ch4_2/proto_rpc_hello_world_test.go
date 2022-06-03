package ch4_2

import (
	"net"
	"net/rpc"
	"testing"

	"github.com/yezihack/colorlog"
)

type HelloService struct{}

func (p *HelloService) Hello(request *String, reply *String) error {
	reply.Value = "hello:" + request.GetValue()
	return nil
}

func TestHelloGeneratedByItherNetRpcPlugin(t *testing.T) {
	server := rpc.NewServer()
	err := RegisterHelloService(server, new(HelloService))
	if err != nil {
		colorlog.Error("register hello service error: %v", err)
	}

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		colorlog.Error("ListenTcp error: %v", err)
	}
	colorlog.Info("Listening...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			colorlog.Error("Accept error: %v", err)
		}

		colorlog.Info("Accepted: %+v", conn)
		go server.ServeConn(conn)
	}
}

func TestClientHelloGeneratedByItherNetRpcPlugin(t *testing.T) {
	client, err := DialHelloService("tcp", ":1234")
	if err != nil {
		colorlog.Error("Dial hello service error: %v", err)
	}

	var reply String
	err = client.Hello(&String{Value: "ither"}, &reply)
	if err != nil {
		colorlog.Error("Call hello service error: %v", err)
	}
	colorlog.Info("reply: %+v", reply)
}
