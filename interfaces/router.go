package interfaces

import "net/http"

type Router struct {
	*http.ServeMux
}

func NewRouter() *Router {
	return &Router{http.NewServeMux()}
}

func (r *Router) ContextHandle(pattern string, handler ContextHandler) {
	r.Handle(pattern, &handler)
}
