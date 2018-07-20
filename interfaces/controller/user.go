package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/koyuta/go-application-sample/domain"
)

type userUsecase interface {
	Add(context.Context, domain.User) error
	Get(context.Context, int64) (domain.User, error)
	GetList(context.Context) ([]domain.User, error)
}

type User struct {
	usecase userUsecase
}

// NewUser returns a new User.
func NewUser(u userUsecase) *User {
	return &User{usecase: u}
}

func (u *User) Get(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	//id, err := strconv.Atoi(r.URL.Path)
	user, err := u.usecase.Get(ctx, 0)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		w.WriteHeader(500)
	}
	b, err := json.Marshal(user)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		w.WriteHeader(500)
	}
	fmt.Fprintf(w, string(b))
}

func (u *User) Post(ctx context.Context, w http.ResponseWriter, r *http.Request) {
}
