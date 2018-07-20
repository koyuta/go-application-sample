package middleware

import (
	"context"
	"net/http"

	"github.com/koyuta/go-application-sample/interfaces"
)

func CommonMiddleware(next interfaces.ContextHandler) interfaces.ContextHandler {
	return func(ctx context.Context, w http.ResponseWriter, r *http.Request) {
		next(ctx, w, r)
	}
}
