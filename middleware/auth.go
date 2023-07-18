package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"net/http"
)

func AuthRequired(jwtSecret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Request.Cookie("session")

		if err != nil {
			if errors.Is(err, http.ErrNoCookie) {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "no session cookie found"})
				c.Abort()
				return
			}

			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get session cookie"})
			c.Abort()
			return
		}

		tokenStr := cookie.Value
		claims := &jwt.MapClaims{}

		_, err = jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (any, error) {
			return []byte(jwtSecret), nil
		})

		if err != nil {
			if errors.Is(err, jwt.ErrSignatureInvalid) {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid session token"})
				c.Abort()
				return
			}

			c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
			c.Abort()
			return
		}

		c.Set("username", (*claims)["username"])
		c.Next()
	}
}
