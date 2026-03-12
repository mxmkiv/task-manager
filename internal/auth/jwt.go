package auth

import (
	"errors"
	"task-manager/internal/model"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(id int, login, role, secretKey string) (string, error) {

	claims := &model.JwtClaims{
		Id:    id,
		Login: login,
		Role:  role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 12)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretKey))
}

func ParseToken(tokenStr, secretKey string) (*model.JwtClaims, error) {
	sc := secretKey

	token, err := jwt.ParseWithClaims(tokenStr, &model.JwtClaims{}, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(sc), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*model.JwtClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token claims")
	}

	return claims, nil
}
