package domain

import "errors"

var (
	ErrCannotCreateInitialUserIfUserTableNotEmpty = errors.New("cannot create initial user if user table is not empty")
)
