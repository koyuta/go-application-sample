package infrastructure

import (
	"context"
	"net/http"

	"github.com/codegangsta/negroni"
)

func Router(ctx context.Context, db *MySQLHandler) *http.ServeMux {
	mux := http.NewServeMux()

	user := controller.NewUser(db)
	mux.Handle("/user", negroni.New(
		negroni.HandlerFunc(middleware.Validate),
		negroni.Wrap(http.HandlerFunc(user.Get)),
	))
	return mux
}
