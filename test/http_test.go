package test

import (
	"log"
	"net/http"
	"testing"
)

type server int

func (h *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Panicln(r.URL.Path)
	w.Write([]byte("hello,world"))
}

func TestHTTP(t *testing.T) {
	var s server
	http.ListenAndServe("localhost:9999", &s)
}
