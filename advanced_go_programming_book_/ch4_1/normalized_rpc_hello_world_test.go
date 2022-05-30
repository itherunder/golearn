package ch4_1

import (
	"net"
	"net/rpc"
	"testing"

	"github.com/yezihack/colorlog"
)

const HelloServiceName = "golearn/advanced_go_programming_book_/ch4_1.HelloService"

type HelloServiceInterface = interface {
	Hello(request string, reply *string) error
}

func RegistrerHelloService(svc HelloServiceInterface) error {
	return rpc.RegisterName(HelloServiceName, svc)
}

func TestNormalizedRpcHelloWorld(t *testing.T) {
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
		go rpc.ServeConn(conn)
	}

	// conn, err := listener.Accept()
	// if err != nil {
	// 	colorlog.Error("Accept error: %v", err)
	// }
	// defer conn.Close()

	// colorlog.Info("Accepted: %+v", conn)

	// rpc.ServeConn(conn)
	// time.Sleep(20 * time.Second)
}

// normalized client
type HelloServiceClient struct {
	*rpc.Client
}

var _ HelloServiceInterface = (*HelloServiceClient)(nil)

func DialHelloService(network, address string) (*HelloServiceClient, error) {
	c, err := rpc.Dial(network, address)
	if err != nil {
		// colorlog.Error("Dial error: %v", err)
		return nil, err
	}
	return &HelloServiceClient{c}, nil
}

func (p *HelloServiceClient) Hello(request string, reply *string) error {
	return p.Client.Call(HelloServiceName+".Hello", request, reply)
}

func TestNormalizedClientRpcRequest(t *testing.T) {
	client, err := DialHelloService("tcp", ":1234")
	if err != nil {
		colorlog.Error("Dial error: %v", err)
	}

	var reply string
	err = client.Hello("ither", &reply)
	if err != nil {
		colorlog.Error("Call rpc service error: %v", err)
	}
	colorlog.Info("reply: %q", reply)
}
