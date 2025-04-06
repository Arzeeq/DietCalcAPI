package auth

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestCreateJWTWithoutError(t *testing.T) {
	login := "some_login"
	jwtParams := JWTParams{
		Secret:   []byte("some_secret"),
		Duration: time.Minute,
	}

	_, err := CreateJWT(login, jwtParams)

	require.NoError(t, err, "JWTCreate must execute without error")
}

func TestGetClaimsJWT(t *testing.T) {
	login := "some_login"
	jwtDuration := time.Minute

	jwtParams := JWTParams{
		Secret:   []byte("some_secret"),
		Duration: jwtDuration,
	}

	token, err := CreateJWT(login, jwtParams)
	require.NoError(t, err, "JWTCreate must execute without error")

	claims, err := GetClaimsJWT(token, jwtParams)
	require.NoError(t, err, "Failed to get claims")

	claimsDuration := claims.ExpiresAt.Time.Sub(claims.IssuedAt.Time)

	require.Equal(t, jwtDuration, claimsDuration, "Incorrect duration")
	require.Equal(t, login, claims.Login, "Incorrect login")
}
