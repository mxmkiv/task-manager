package encoder

import "golang.org/x/crypto/bcrypt"

type HashEncoder interface {
	Encode(password string) (string, error)
	Compare(hash, password string) error
}

type BcryptEncoder struct{}

func NewBcryptEncoder() HashEncoder {
	return &BcryptEncoder{}
}

func (u *BcryptEncoder) Encode(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func (u *BcryptEncoder) Compare(hash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
