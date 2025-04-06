package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTClaims struct {
	jwt.RegisteredClaims
	Login string
}

type JWTParams struct {
	Secret   []byte
	Duration time.Duration
}

func CreateJWT(login string, params JWTParams) (string, error) {
	now := time.Now()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, JWTClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(params.Duration)),
		},
		Login: login,
	})

	signedToken, err := token.SignedString(params.Secret)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func GetClaimsJWT(token string, params JWTParams) (*JWTClaims, error) {
	keyFunc := func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return params.Secret, nil
	}

	claims := &JWTClaims{}
	parsedToken, err := jwt.ParseWithClaims(token, claims, keyFunc)
	if err != nil {
		return nil, fmt.Errorf("token parsing error: %v", err)
	}

	if !parsedToken.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}
