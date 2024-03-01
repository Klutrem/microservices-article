package domainCommon

import (
	"github.com/golang-jwt/jwt/v5"
)

type JwtCustomClaims struct {
	Name string `json:"name"`
	ID   string `json:"id"`
	jwt.Claims
}

type JwtCustomRefreshClaims struct {
	ID string `json:"id"`
	jwt.Claims
}
