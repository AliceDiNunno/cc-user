package postgres

import (
	"github.com/AliceDiNunno/cc-user/src/core/domain"
	e "github.com/AliceDiNunno/go-nested-traced-error"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type userTokenRepo struct {
	db *gorm.DB
}

type UserToken struct {
	gorm.Model
	ID     uuid.UUID `gorm:"type:uuid;primary_key"`
	Token  string
	Name   string
	UserId uuid.UUID
	User   *User
}

func (u userTokenRepo) CreateToken(token *domain.UserToken) *e.Error {
	userTokenToCreate := userTokenFromDomain(token)

	result := u.db.Create(userTokenToCreate)

	if result.Error != nil {
		return e.Wrap(result.Error)
	}

	return nil
}

func userTokenToDomain(user *UserToken) *domain.UserToken {
	return &domain.UserToken{
		ID:    user.ID,
		Token: user.Token,
		Name:  user.Name,
		User:  userToDomain(user.User),
	}
}

func userTokenFromDomain(user *domain.UserToken) *UserToken {
	return &UserToken{
		ID:     user.ID,
		Token:  user.Token,
		Name:   user.Name,
		UserId: user.ID,
		User:   userFromDomain(user.User),
	}
}

func NewUserTokenRepo(db *gorm.DB) userTokenRepo {
	return userTokenRepo{
		db: db,
	}
}
