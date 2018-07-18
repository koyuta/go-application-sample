package middleware

import (
	"net/http"
	"strconv"
)

func Validate(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	p := r.URL.Query()
	if m, err := strconv.Atoi(p.Get("name")); m < 1 || err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	next(w, r)
}
