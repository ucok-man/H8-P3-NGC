package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTClaim struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func NewJWTClaim(username string, expiration time.Time) JWTClaim {
	return JWTClaim{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiration),
		},
	}
}

func GenerateToken(claim *JWTClaim, secret string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString([]byte(secret))
}

func DecodeToken(inputToken string, claim *JWTClaim, secret string) error {
	token, err := jwt.ParseWithClaims(inputToken, claim, func(token *jwt.Token) (any, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return err
	}

	if !token.Valid {
		return err
	}
	return nil
}
