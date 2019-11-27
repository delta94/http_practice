package main

import (
	"fmt"
	"github.com/parkjinhong03/http_practice/net"
	"net/http"
)

func main() {
	s := net.NewServer()

	s.HandleFunc("GET", "/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "asdf")
	})
	s.HandleFunc("GET", "/1", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "asdf")
	})
	s.HandleFunc("GET", "/12", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "asdf")
	})
	s.HandleFunc("GET", "/123", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "asdf")
	})

	s.Run(":8090")
}
