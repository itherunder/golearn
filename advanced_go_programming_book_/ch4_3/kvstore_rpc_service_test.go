package ch4_3

import (
	"fmt"
	"math/rand"
	"net"
	"net/rpc"
	"sync"
	"testing"
	"time"

	"github.com/yezihack/colorlog"
)

type KVStoreService struct {
	m      map[string]string
	filter map[string]func(key string)
	mu     sync.RWMutex
}

func NewKVStoreService() *KVStoreService {
	return &KVStoreService{
		m:      make(map[string]string),
		filter: make(map[string]func(key string)),
	}
}

func (p *KVStoreService) Set(kv [2]string, reply *struct{}) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	key, value := kv[0], kv[1]

	if oldValue := p.m[key]; oldValue != value {
		for _, fn := range p.filter {
			fn(key)
		}
	}
	p.m[key] = value
	return nil
}

func (p *KVStoreService) Get(key string, reply *string) error {
	p.mu.RLock()
	defer p.mu.RUnlock()

	if v, ok := p.m[key]; ok {
		*reply = v
		return nil
	}

	return fmt.Errorf("not found")
}

func (p *KVStoreService) Watch(timeout time.Duration, keyChanged *string) error {
	id := fmt.Sprintf("watch-%s-%03d", time.Now(), rand.Int())
	ch := make(chan string, 10) // bufferred

	p.mu.Lock()
	p.filter[id] = func(key string) { ch <- key }
	p.mu.Unlock()

	select {
	case <-time.After(timeout):
		return fmt.Errorf("timeout")
	case v := <-ch:
		*keyChanged = v
		return nil
	}
}

func TestKVStoreRpcService(t *testing.T) {
	rpc.RegisterName("KVStoreService", NewKVStoreService())

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
}

func TestDoClientKVStoreRpc(t *testing.T) {
	client, err := rpc.Dial("tcp", ":1234")
	if err != nil {
		colorlog.Error("dial kvstore service error: %v", err)
	}

	go func() {
		var keyChanged string
		err := client.Call("KVStoreService.Watch", 3*time.Second, &keyChanged)
		if err != nil {
			colorlog.Error("call kvstore service error: %v", err)
		}
		colorlog.Info("result: %q", keyChanged)
	}()
	time.Sleep(time.Second)

	err = client.Call("KVStoreService.Set", [2]string{"ither", "under"}, new(struct{}))
	if err != nil {
		colorlog.Error("call kvstore service error: %v", err)
	}
	time.Sleep(3 * time.Second)
}
