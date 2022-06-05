package ch4_3

import (
	"golearn/advanced_go_programming_book_/ch4_2"
	"testing"

	"github.com/yezihack/colorlog"
)

func TestDoClientRpc(t *testing.T) {
	client, err := ch4_2.DialHelloService("tcp", ":1234")
	if err != nil {
		colorlog.Error("dial hello service error: %v", err)
	}
	helloCall := client.Go("HelloService.Hello", ch4_2.String{Value: "ither"}, new(ch4_2.String), nil)
	colorlog.Info("waiting for reply...")
	helloCall = <-helloCall.Done
	if err := helloCall.Error; err != nil {
		colorlog.Error("rpc service hello call error: %v", err)
	}

	args := helloCall.Args.(ch4_2.String)
	reply := helloCall.Reply.(*ch4_2.String)
	colorlog.Info("args: %+v, reply: %+v", args, reply)
}
