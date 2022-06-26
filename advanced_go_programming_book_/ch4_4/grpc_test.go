package ch4_4

import (
	"context"
	"net"
	"testing"

	"github.com/yezihack/colorlog"
	grpc "google.golang.org/grpc"
)

type HelloServiceImpl struct{}

func (p *HelloServiceImpl) Hello(ctx context.Context, in *String) (*String, error) {
	reply := &String{
		Value: "hello:" + in.GetValue(),
	}
	return reply, nil
}

func TestGrpc(t *testing.T) {
	grpcServer := grpc.NewServer()
	RegisterHelloServiceServer(grpcServer, new(HelloServiceImpl))

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		colorlog.Error("listen error: %v", err)
	}
	err = grpcServer.Serve(listener)
	if err != nil {
		colorlog.Error("serve error: %v", err)
	}
}

func TestGrpcClient(t *testing.T) {
	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
	if err != nil {
		colorlog.Error("dial error: %v", err)
	}
	client := NewHelloServiceClient(conn)
	reply, err := client.Hello(context.Background(), &String{Value: "ither"})
	if err != nil {
		colorlog.Error("hello error: %v", err)
	}
	colorlog.Info("reply is %q", reply.GetValue())
}
