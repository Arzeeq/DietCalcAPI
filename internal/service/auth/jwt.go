package auth

import (
	"dietcalc/internal/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateJWT(secret []byte, login string) (string, error) {
	jwtDuration := config.Cfg.JWTDuration

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"login":      login,
		"expires_at": time.Now().Add(jwtDuration).Unix(),
	})

	signedToken, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
