package util

import (
	"github.com/ermites-io/passwd"
)

// Hasher password hashing service
type Hasher struct {
	p *passwd.Profile
}

func NewHash() (*Hasher, error) {
	p, err := passwd.New(passwd.Argon2idDefault)
	if err != nil {
		return nil, err
	}

	return &Hasher{p: p}, nil
}

func (h *Hasher) Hash(psw string) (string, error) {
	hashed, err := h.p.Hash([]byte(psw))
	if err != nil {
		return "", err
	}

	return string(hashed), nil
}

func (h *Hasher) Compare(psw string, hashedPsw string) (bool, error) {
	err := passwd.Compare([]byte(hashedPsw), []byte(psw))
	if err != nil {
		if err == passwd.ErrMismatch {
			return false, nil
		}

		return false, err
	}

	return true, nil
}
