package usecases

import (
	"github.com/AliceDiNunno/cc-user/src/config"
	"github.com/AliceDiNunno/cc-user/src/core/domain"
	e "github.com/AliceDiNunno/go-nested-traced-error"
)

type Usecases interface {
	//Authentication
	CreateAuthToken(request domain.AccessTokenRequest) (string, *e.Error)
	CreateJwtToken(request domain.JwtTokenRequest) (string, *e.Error)
	CheckJwtToken(token string) (*domain.JwtTokenPayload, *e.Error)

	CreateInitialUser(user *config.InitialUserConfig) *e.Error
	CreateUser(user *domain.UserCreationRequest) *e.Error
}
