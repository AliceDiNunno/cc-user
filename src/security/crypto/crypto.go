package crypto

import "golang.org/x/crypto/bcrypt"

func HashAndSalt(token string) (string, error) {
	//TODO: cost should be a variable
	hash, err := bcrypt.GenerateFromPassword([]byte(token), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
