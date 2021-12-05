package domain

import (
	"github.com/google/uuid"
)

type UserToken struct {
	ID   uuid.UUID
	User *User

	Token string
	Name  string
}

func (u *UserToken) Initialize() {
	u.ID = uuid.New()
}
