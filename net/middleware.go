package net

import (
	"log"
	"net/http"
	"time"
)

type middleware func(next HandlerFunc) HandlerFunc

func LogHandler (next HandlerFunc) HandlerFunc {
	return func(c *Context) {
		t := time.Now()

		next(c)

		log.Printf("[%s] %q %v\n",
			c.Request.Method,
			c.Request.URL.String(),
			time.Now().Sub(t))
	}
}

func RecoverHandler (next HandlerFunc) HandlerFunc {
	return func(c *Context) {
		defer func() {
			if err:=recover(); err!=nil {
				log.Printf("panic: %+v", err)
				http.Error(c.ResponseWriter,
					http.StatusText(http.StatusInternalServerError),
					http.StatusInternalServerError)
			}
		}()
		next(c)
		return
	}
}