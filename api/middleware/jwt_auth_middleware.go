package middleware

import (
	"net/http"
	"strings"

	"main/internal/tokenutil"

	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware(publicKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		t := strings.Split(authHeader, " ")
		if len(t) == 2 {
			authToken := t[1]
			authorized, err := tokenutil.IsAuthorized(authToken, publicKey)
			if authorized {
				userID, err := tokenutil.ExtractIDFromToken(authToken, publicKey)
				if err != nil {
					c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
					c.Abort()
					return
				}
				c.Set("account_id", userID)
				c.Next()
				return
			}
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unathorized"})
		c.Abort()
	}
}
