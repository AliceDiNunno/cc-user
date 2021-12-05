package postgres

import (
	"github.com/AliceDiNunno/cc-user/src/core/domain"
	e "github.com/AliceDiNunno/go-nested-traced-error"
	"gorm.io/gorm"
)
import "github.com/google/uuid"

type User struct {
	gorm.Model
	ID       uuid.UUID `gorm:"type:uuid;primary_key"`
	Mail     string
	Password string
}

func (u userRepo) IsEmpty() bool {
	var count int64

	u.db.Model(&User{}).Count(&count)

	return count == 0
}

func (u userRepo) CreateUser(user *domain.User) *e.Error {
	userToCreate := userFromDomain(user)

	result := u.db.Create(userToCreate)

	if result.Error != nil {
		return e.Wrap(result.Error)
	}

	return nil
}

type userRepo struct {
	db *gorm.DB
}

func userToDomain(user *User) *domain.User {
	return &domain.User{
		ID:       user.ID,
		Mail:     user.Mail,
		Password: user.Password,
	}
}

func userFromDomain(user *domain.User) *User {
	return &User{
		ID:       user.ID,
		Mail:     user.Mail,
		Password: user.Password,
	}
}

func NewUserRepo(db *gorm.DB) userRepo {
	return userRepo{
		db: db,
	}
}
