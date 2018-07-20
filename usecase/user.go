package usecase

import (
	"context"

	"github.com/koyuta/go-application-sample/domain"
)

type userRepository interface {
	Store(context.Context, domain.User) (int64, error)
	FindByID(context.Context, int64) (domain.User, error)
	FindAll(context.Context) ([]domain.User, error)
}

type User struct {
	repository userRepository
}

func NewUser(r userRepository) *User {
	return &User{repository: r}
}

func (u *User) Add(ctx context.Context, user domain.User) error {
	_, err := u.repository.Store(ctx, user)
	return err
}

func (u *User) Get(ctx context.Context, id int64) (domain.User, error) {
	return u.repository.FindByID(ctx, id)
}

func (u *User) GetList(ctx context.Context) ([]domain.User, error) {
	return u.repository.FindAll(ctx)
}
