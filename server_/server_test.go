package server_

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"testing"

	"github.com/yezihack/colorlog"
)

var mu sync.Mutex
var count int

func TestServer(t *testing.T) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("<h1>home</h1>"))
	})
	http.HandleFunc("/:<page>", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("<h1>hello</h1>"))
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func TestServerCount(t *testing.T) {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}

func TestEchoServer(t *testing.T) {
	http.HandleFunc("/echo", echo)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func echo(w http.ResponseWriter, r *http.Request) {
	colorlog.Info("%s %s %s", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		colorlog.Info("Header[%q] = %q", k, v)
	}
	colorlog.Info("Host = %q", r.Host)
	colorlog.Info("RemoteAddr = %q", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		colorlog.Error("ParseForm() err: %v", err)
	}
	for k, v := range r.Form {
		colorlog.Info("Form[%q] = %q", k, v)
	}
	w.Write([]byte("<h1>" + r.Form.Get("msg") + "</h1>"))
}
