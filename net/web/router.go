package web

import (
	"net/http"
)

// web http pattern router
type Router struct {
	mux     *http.ServeMux
	pattern string
}

// NewRouter new a router.
func NewRouter(mux *http.ServeMux) *Router {
	return &Router{mux: mux}
}

func (r *Router) join(pattern string) string {
	return r.pattern + pattern
}

func (r *Router) Group(pattern string) *Router {
	return &Router{mux: r.mux, pattern: r.join(pattern)}
}

// Handler is an adapter which allows the usage of an http.Handler as a
// request handle.
func (r *Router) Handle(method, pattern string, handlers ...Handler) {
	r.mux.HandleFunc(r.join(pattern), func(w http.ResponseWriter, r *http.Request) {
		handler(method, w, r, handlers)
	})
}

func (r *Router) HandleFunc(method, pattern string, handlers ...HandleFunc) {
	r.mux.HandleFunc(r.join(pattern), func(w http.ResponseWriter, r *http.Request) {
		handleFunc(method, w, r, handlers)
	})
}

// Get is a shortcut for router.Handle("GET", path, handle)
func (r *Router) Get(pattern string, handlers ...Handler) {
	r.mux.HandleFunc(r.join(pattern), func(w http.ResponseWriter, r *http.Request) {
		handler("GET", w, r, handlers)
	})
}

func (r *Router) Post(pattern string, handlers ...Handler) {
	r.mux.HandleFunc(r.join(pattern), func(w http.ResponseWriter, r *http.Request) {
		handler("POST", w, r, handlers)
	})
}

// GetFunc is a shortcut for router.HandleFunc("GET", path, handle)
func (r *Router) GetFunc(pattern string, handlers ...HandleFunc) {
	r.mux.HandleFunc(r.join(pattern), func(w http.ResponseWriter, r *http.Request) {
		handleFunc("GET", w, r, handlers)
	})
}

// PostFunc is a shortcut for router.HandleFunc("GET", path, handle)
func (r *Router) PostFunc(pattern string, handlers ...HandleFunc) {
	r.mux.HandleFunc(r.join(pattern), func(w http.ResponseWriter, r *http.Request) {
		handleFunc("POST", w, r, handlers)
	})
}

func handler(method string, w http.ResponseWriter, r *http.Request, handlers []Handler) {
	if r.Method != method {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	// TODO reuse context
	c := new(Context)
	c.Request = r
	c.Response = w
	for _, h := range handlers {
		h.ServeHTTP(c)
		if c.done {
			break
		}
	}
}

func handleFunc(method string, w http.ResponseWriter, r *http.Request, handlers []HandleFunc) {
	if r.Method != method {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	// TODO reuse context
	c := new(Context)
	c.Request = r
	c.Response = w
	for _, h := range handlers {
		h(c)
		if c.done {
			break
		}
	}
}
