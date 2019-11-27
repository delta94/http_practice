package net

import (
	"log"
	"net/http"
	"time"
)

type middleware func(http.HandlerFunc) http.HandlerFunc

func LogHandler (next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()

		next(w, r)

		log.Printf("[%s] %q %v\n",
			r.Method,
			r.URL.String(),
			time.Now().Sub(t))
	}
}

func RecoverHandler (next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err:=recover(); err!=nil {
				log.Printf("panic: %+v", err)
				http.Error(w,
					http.StatusText(http.StatusInternalServerError),
					http.StatusInternalServerError)
			}
		}()
		next(w, r)
		return
	}
}