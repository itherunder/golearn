package ch4_3

import (
	"net"
	"net/rpc"
	"testing"
	"time"
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
		go func() {
			rpc.ServeConn(conn)
			conn.Close()
		}()
	}
}
