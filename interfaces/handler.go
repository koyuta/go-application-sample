package interfaces

import (
	"context"
	"net/http"
	"time"
)

var Timeout = 50 * time.Millisecond

type ContextHandler func(context.Context, http.ResponseWriter, *http.Request)

func (c ContextHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), Timeout)
	defer cancel()
	c(ctx, w, r)
}
