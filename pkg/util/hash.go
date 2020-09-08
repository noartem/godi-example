package util

import "golang.org/x/crypto/bcrypt"

func HashPassword(psw string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(psw), bcrypt.MaxCost)
	if err != nil {
		return "", err
	}

	return string(hashed), nil
}

func CompareHash(psw string, hashedPsw string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(psw), []byte(hashedPsw))
	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return false, nil
		}

		return false, err
	}

	return true, nil
}
