package domain

import "github.com/google/uuid"

type UserCreationRequest struct {
	Email    string
	Password string
}

type User struct {
	ID       uuid.UUID
	Mail     string
	Password string
}
