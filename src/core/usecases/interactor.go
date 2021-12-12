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
	FindByMail(mail string) (*domain.User, *e.Error)
}

type UserTokenRepo interface {
	CreateToken(token *domain.AccessToken) *e.Error
	FindByToken(token string) (*domain.AccessToken, *e.Error)
}

type JwtSignatureRepo interface {
	SaveSignature(signature *domain.JwtSignature) *e.Error
	CheckIfSignatureExists(signature string) bool
}

type interactor struct {
	userRepo      UserRepo
	userTokenRepo UserTokenRepo
	jwtSignature  JwtSignatureRepo
}

func NewInteractor(u UserRepo, ut UserTokenRepo, js JwtSignatureRepo) interactor {
	return interactor{
		userRepo:      u,
		userTokenRepo: ut,
		jwtSignature:  js,
	}
}
