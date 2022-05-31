package ch4_1

import (
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"testing"

	"github.com/yezihack/colorlog"
)

func TestJsonRpcHelloWorld(t *testing.T) {
	// err := RegistrerHelloService(&HelloService{})
	err := RegistrerHelloService(new(HelloService))
	if err != nil {
		colorlog.Error("Register hello service error: %v", err)
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
		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}

type JsonHelloServiceClient struct {
	*rpc.Client
}

func DialJsonHelloService(network, address string) (*JsonHelloServiceClient, error) {
	conn, err := net.Dial(network, address)
	if err != nil {
		colorlog.Error("Dial hello service error: %v", err)
		return nil, err
	}
	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))
	return &JsonHelloServiceClient{client}, nil
}

func (p *JsonHelloServiceClient) Hello(request string, reply *string) error {
	return p.Client.Call(HelloServiceName+".Hello", request, reply)
}

func TestClientJsonRpcRequest(t *testing.T) {
	client, err := DialJsonHelloService("tcp", ":1234")
	if err != nil {
		colorlog.Error("Dial hello service error: %v", err)
	}

	var reply string
	err = client.Hello("ither", &reply)
	if err != nil {
		colorlog.Error("Call hello service error: %v", err)
	}
	colorlog.Info("reply: %q", reply)
}
