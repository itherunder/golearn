package practice_

import (
	"fmt"
	"strings"
	"sync"
	"testing"
	"time"
)

type (
	subscriber chan interface{}
	topicFunc  func(v interface{}) bool
)

type Publisher struct {
	m           sync.RWMutex
	buffer      int
	timeout     time.Duration
	subscribers map[subscriber]topicFunc
}

func NewPublisher(publishTimeout time.Duration, buffer int) *Publisher {
	return &Publisher{
		timeout:     publishTimeout,
		buffer:      buffer,
		subscribers: make(map[subscriber]topicFunc),
	}
}

func (p *Publisher) Subscribe() subscriber {
	return p.SubscribeTopic(nil)
}

func (p *Publisher) SubscribeTopic(topic topicFunc) subscriber {
	p.m.Lock()
	defer p.m.Unlock()
	ch := make(chan interface{}, p.buffer)
	p.subscribers[ch] = topic
	return ch
}

func (p *Publisher) Evict(sub subscriber) {
	p.m.Lock()
	defer p.m.Unlock()
	delete(p.subscribers, sub)
	close(sub)
}

func (p *Publisher) Publish(v interface{}) {
	p.m.RLock()
	defer p.m.RUnlock()

	var wg sync.WaitGroup
	for sub, topic := range p.subscribers {
		wg.Add(1)
		go p.sendTopic(sub, topic, v, &wg)
	}

	wg.Wait()
}

func (p *Publisher) Close() {
	p.m.Lock()
	defer p.m.Unlock()

	for sub := range p.subscribers {
		delete(p.subscribers, sub)
		close(sub)
	}
}

func (p *Publisher) sendTopic(sub subscriber, topic topicFunc, v interface{}, wg *sync.WaitGroup) {
	defer wg.Done()
	if topic != nil && !topic(v) {
		return
	}
	select {
	case sub <- v:
	case <-time.After(p.timeout):
	}
}

func TestPubSub(t *testing.T) {
	p := NewPublisher(5*time.Second, 32)
	defer p.Close()

	all := p.Subscribe()
	golang := p.SubscribeTopic(func(v interface{}) bool {
		if s, ok := v.(string); ok {
			return strings.Contains(s, "golang")
		}
		return false
	})

	p.Publish("hello, world!")
	p.Publish("hello, golang!")

	go func() {
		for msg := range all {
			fmt.Println("all: ", msg)
		}
	}()

	go func() {
		for msg := range golang {
			fmt.Println("golang: ", msg)
		}
	}()

	time.Sleep(5 * time.Second)
}
