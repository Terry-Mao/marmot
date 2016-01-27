package web

type Handler interface {
	ServeHTTP(*Context)
}

type HandleFunc func(*Context)
