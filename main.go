package main

import (
	"fmt"
	"github.com/parkjinhong03/http_practice/net"
)

func CustomMiddleware(next net.HandlerFunc) net.HandlerFunc {
	return func(c *net.Context) {
		fmt.Println(c.Request.Method, "요청이 들어왔습니다!")

		next(c)
		return
	}
}

func main() {
	s := net.NewServer()

	s.HandleFunc("GET", "/:user_id", func(c *net.Context) {
		fmt.Fprintln(c.ResponseWriter, "asdf", c.Params["user_id"])
	})
	s.HandleFunc("GET", "/1", func(c *net.Context) {
		fmt.Fprintln(c.ResponseWriter, "asdf")
	})
	s.HandleFunc("GET", "/12", func(c *net.Context) {
		fmt.Fprintln(c.ResponseWriter, "asdf")
	})
	s.HandleFunc("GET", "/123", func(c *net.Context) {
		fmt.Fprintln(c.ResponseWriter, "asdf")
	})

	s.Use(CustomMiddleware)

	s.Run(":8090")
}
