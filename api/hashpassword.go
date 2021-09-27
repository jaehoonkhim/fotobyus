package api

import (
	"golang.org/x/crypto/bcrypt"
)

type HashPassword struct {
	pepper string
}

func NewHashPassword(in string) *HashPassword {
	return &HashPassword{pepper: in}
}

func (h *HashPassword) Generate(plain string) (string, error) {
	hp, err := bcrypt.GenerateFromPassword([]byte(h.pepper+plain), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hp), nil
}

func (h *HashPassword) Compare(hash, plain string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(h.pepper+plain))
}
