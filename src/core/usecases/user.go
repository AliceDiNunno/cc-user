package usecases

import (
	"github.com/AliceDiNunno/cc-user/src/config"
	"github.com/AliceDiNunno/cc-user/src/core/domain"
	e "github.com/AliceDiNunno/go-nested-traced-error"
)

func (i interactor) CreateUser(user *domain.UserCreationRequest) e.Error {
	panic("implement me")
}

func (i interactor) CreateInitialUser(user *config.InitialUserConfig) *e.Error {
	panic("implement me")
}
