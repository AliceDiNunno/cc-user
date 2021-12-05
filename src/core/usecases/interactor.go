package usecases

import (
	"github.com/AliceDiNunno/cc-user/src/core/domain"
	e "github.com/AliceDiNunno/go-nested-traced-error"
)

type Logger interface {
	Error(args ...interface{})
	Fatal(args ...interface{})
	Info(args ...interface{})
	Debug(args ...interface{})
}

type UserRepo interface {
	IsEmpty() bool
	CreateUser(user *domain.User) *e.Error
}

type interactor struct {
	userRepo UserRepo
}

func NewInteractor(u UserRepo) interactor {
	return interactor{
		userRepo: u,
	}
}
