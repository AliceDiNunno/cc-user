package usecases

import (
	"github.com/AliceDiNunno/cc-user/src/config"
	"github.com/AliceDiNunno/cc-user/src/core/domain"
	"github.com/AliceDiNunno/cc-user/src/security/crypto"
	e "github.com/AliceDiNunno/go-nested-traced-error"
)

func (i interactor) CreateUser(user *domain.UserCreationRequest) *e.Error {
	panic("implement me")
}

func (i interactor) CreateInitialUser(user *config.InitialUserConfig) *e.Error {
	if !i.userRepo.IsEmpty() {
		return e.Wrap(domain.ErrCannotCreateInitialUserIfUserTableNotEmpty)
	}

	hash, stderr := crypto.HashAndSalt(user.Password)

	if stderr != nil {
		return e.Wrap(stderr)
	}

	userToCreate := &domain.User{
		Mail:     user.Mail,
		Password: hash,
	}

	userToCreate.Initialize()

	err := i.userRepo.CreateUser(userToCreate)

	if err != nil {
		return err
	}

	return nil
}
