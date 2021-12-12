package domain

import (
	"bytes"
	"crypto/rand"
	"github.com/google/uuid"
	"math/big"
	"time"
)

type AccessTokenRequest struct {
	Mail     string
	Password string
	OtpCode  string
}

type AccessToken struct {
	CreatedAt         time.Time
	ID                uuid.UUID
	User              *User
	Token             string
	IsPersonnalAccess bool
	JwtGenerated      []*JwtSignature
}

func (t *AccessToken) Valid() bool {
	if t == nil {
		println("UserToken Nil")
		return false
	}

	if t.IsPersonnalAccess {
		println("UserToken is personnal")
		return true
	}

	if time.Now().Before(t.CreatedAt) {
		println("UserToken is created before now")
		return false
	}

	if t.CreatedAt.Add(time.Minute * 30).Before(time.Now()) {
		println("UserToken is expired")
		return false
	}
	println("UserToken is valid")
	return true
}

func generateRandomString(n int) (string, error) {
	const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-"
	ret := make([]byte, n)
	for i := 0; i < n; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			return "", err
		}
		ret = append(ret, letters[num.Int64()])
	}

	ret = bytes.Trim(ret, "\x00")
	return string(ret), nil
}

func (u *AccessToken) Initialize() {
	u.ID = uuid.New()

	if u.Token == "" {
		token, err := generateRandomString(32)
		if err == nil {
			u.Token = token
		}
	}
}
