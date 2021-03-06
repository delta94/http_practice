package net

import "net/http"

type Context struct {
	Params map[string]interface{}
	ResponseWriter http.ResponseWriter
	Request *http.Request
}

type HandlerFunc func(c *Context)