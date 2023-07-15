package tokenutil

import (
	"errors"

	"github.com/golang-jwt/jwt/v5"
)

func IsAuthorized(requestToken string, secret string) (bool, error) {
	key, err := jwt.ParseRSAPublicKeyFromPEM([]byte(secret))
	token, err := jwt.Parse(requestToken, func(t *jwt.Token) (interface{}, error) {
		if err != nil {
			return false, err
		}
		return key, nil
	})
	if err != nil {
		return false, err
	}
	if token.Valid {
		return true, nil

	}
	return false, errors.New("token invalid")
}

func ExtractIDFromToken(requestToken string, secret string) (string, error) {
	key, err := jwt.ParseRSAPublicKeyFromPEM([]byte(secret))
	token, err := jwt.Parse(requestToken, func(t *jwt.Token) (interface{}, error) {
		if err != nil {
			return false, errors.New("token invalid")
		}
		return key, nil
	})
	if err != nil {
		return "", err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		id := claims["sub"].(string)
		return id, nil
	}

	return "", errors.New("invalid token")
}
