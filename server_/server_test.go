package server_

import (
	"log"
	"net/http"
	"testing"
)

func TestServer(t *testing.T) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("<h1>home</h1>"))
	})
	http.HandleFunc("/:<page>", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("<h1>{page}</h1>"))
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
