package controller

import (
	"context"
	"fmt"
	"net/http"
	"os"
)

type User struct {
	u *usecase.User
}

func NewUser(db repository.mysql) *User {
	return &User{
		u: usecase.NewUser(repository.User{db}),
	}
}

func (u *User) Get(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	user, err := u.Get(ctx, id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		w.WriteHeader(500)
		return
	}
}

func (u *User) Post(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
}
