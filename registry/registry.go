package registry

import (
	"github.com/koyuta/go-application-sample/interfaces/controller"
	"github.com/koyuta/go-application-sample/interfaces/repository"
	"github.com/koyuta/go-application-sample/usecase"
)

type UserRegistry struct {
	datastore repository.MySQLHandler
}

func NewUserRegistry(datastore repository.MySQLHandler) *UserRegistry {
	return &UserRegistry{datastore: datastore}
}

func (u *UserRegistry) Registry() *controller.User {
	return controller.NewUser(usecase.NewUser(&repository.User{u.datastore}))
}
