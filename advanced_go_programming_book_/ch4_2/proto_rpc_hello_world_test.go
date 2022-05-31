package ch4_2

import (
	"testing"

	"github.com/yezihack/colorlog"
)

type HelloService struct{}

func (p *HelloService) Hello(request *String, reply *String) error {
	reply.Value = "hello:" + request.GetValue()
	return nil
}

func TestHello(t *testing.T) {
	colorlog.Info("test hello")
}
