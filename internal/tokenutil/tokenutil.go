package tokenutil

import (
	"context"
	"errors"
	"fmt"

	"github.com/Nerzal/gocloak/v13"
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
	// secrett, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(os.Getenv("CLIENT_SECRET")))
	// token, err := jwt.Parse(requestToken, func(token *jwt.Token) (interface{}, error) {
	// 	if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
	// 		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	// 	}
	// 	return secrett, nil
	// })

	client := gocloak.NewClient("http://localhost:8080")

	token, claims, err := client.DecodeAccessToken(context.TODO(), requestToken, "aura")
	if err != nil {
		return "", err
	}

	if !token.Valid || claims.Valid() != nil {
		return "", errors.New("invalid token")
	}
	info, err := client.RetrospectToken(context.Background(), requestToken, "admin-cli", secret, "aura")
	fmt.Println(info)
	return "nil", nil
	// claims, ok := claims.(jwt.Claims)

	// fmt.Println(claims.GetSubject())

	// if !ok && !token.Valid {
	// 	return "", fmt.Errorf("invalid Token")
	// }

	// return claims["id"].(string), nil
}
