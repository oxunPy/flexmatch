package security

import (
	"github.com/golang-jwt/jwt/v5"
)

func NewToken(claims *jwt.MapClaims, secret string) (*string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, *claims)
	tokenStr, err := token.SignedString([]byte(secret))
	if err != nil {
		return nil, err
	}

	return &tokenStr, nil
}
