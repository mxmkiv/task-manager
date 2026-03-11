package auth

import "golang.org/x/crypto/bcrypt"

type HashEncoder interface {
	Encode(password string) (string, error)
	Compare(hash, password string) error
}

type UserEncoder struct{}

func NewUserEncoder() HashEncoder {
	return &UserEncoder{}
}

func (u *UserEncoder) Encode(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func (u *UserEncoder) Compare(hash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
