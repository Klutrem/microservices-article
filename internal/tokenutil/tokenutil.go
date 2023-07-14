package tokenutil

import (
	"context"
	"errors"
	"fmt"

	"github.com/Nerzal/gocloak/v13"
	"github.com/golang-jwt/jwt/v5"
)

func IsAuthorized(requestToken string, secret string) (bool, error) {
	// secrett, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(os.Getenv("CLIENT_SECRET")))
	// _, err = jwt.Parse(requestToken, func(token *jwt.Token) (interface{}, error) {
	// 	if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
	// 		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	// 	}
	// 	return secrett, nil
	// })
	// if err != nil {
	// 	return false, err
	// }
	// return true, nil

	client := gocloak.NewClient("http://localhost:8080")

	token, claims, err := client.DecodeAccessToken(context.TODO(), requestToken, "aura")
	if err != nil {
		return false, err
	}

	if !token.Valid || claims.Valid() != nil {
		return false, errors.New("invalid token")
	}

	return true, nil
}

func ExtractIDFromToken(requestToken string, secret string) (string, error) {
	// keyData, err := ioutil.ReadFile("public.key")

	key, err := jwt.ParseRSAPublicKeyFromPEM([]byte(secret))

	// key, err := jwt.ParseRSAPublicKeyFromPEM(keyData)
	token, err := jwt.Parse(requestToken, func(t *jwt.Token) (interface{}, error) {
		if err != nil {
			return false, errors.New("token invalid")
		}
		return key, nil
	})
	if err != nil {
		return "", err
	}
	fmt.Println(token.Claims)

	return "", nil
}
