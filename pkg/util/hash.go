package util

import "golang.org/x/crypto/bcrypt"

func Hash(psw string) (string, error) {
	hashed, err :=  bcrypt.GenerateFromPassword([]byte(psw), bcrypt.MaxCost)
	if err != nil {
		return "", err
	}

	return string(hashed), nil
}
