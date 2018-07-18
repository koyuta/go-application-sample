package usecase

import "context"

type userRepository interface {
	Store(context.Context, domain.User) (int64, error)
	FindByID(int64) (context.Context, domain.User, error)
	FindAll(context.Context) ([]domain.User, error)
}

type User struct {
	repository userRepository
}

func NewUser(r userRepository) *User {
	return &User{repository: r}
}

func (u *User) Add(ctx context.Context, u domain.User) error {
	_, err := u.Repository.Store(ctx, u)
	return err
}

func (u *User) Get(ctx context.Context, id int64) (domain.User, error) {
	return u.Repository.FindByID(ctx, id)
}

func (u *User) GetList(ctx context.Context) ([]domain.User, error) {
	return u.Repository.FindAll(ctx)
}
