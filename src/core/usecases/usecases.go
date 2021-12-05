package usecases

import (
	"github.com/AliceDiNunno/cc-user/src/config"
	"github.com/AliceDiNunno/cc-user/src/core/domain"
	e "github.com/AliceDiNunno/go-nested-traced-error"
)

type Usecases interface {
	CreateInitialUser(user *config.InitialUserConfig) *e.Error
	CreateUser(user *domain.UserCreationRequest) *e.Error
}
