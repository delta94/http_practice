package net

import "net/http"

type server struct {
	*router
	middlewares []middleware
	startHandler http.HandlerFunc
}

func NewServer() *server {
	r := &router{map[string]map[string]http.HandlerFunc{}}
	s := &server{router:r}
	s.middlewares = []middleware{
		// ... 미들 웨어 추가
		LogHandler,
		RecoverHandler,
	}
	return s
}

func (s *server) Run(addr string) {
	s.startHandler = s.router.handler()

	for i := len(s.middlewares)-1; i>=0; i-- {
		s.startHandler = s.middlewares[i](s.startHandler)
	}

	if err := http.ListenAndServe(addr, s); err != nil {
		panic(err)
	}
}

func (s *server) Use (middleware... middleware) {
	s.middlewares = append(s.middlewares, middleware...)
}

func (s *server) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	s.startHandler(w, req)
}
