package middleware

import (
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"tspo_final/internal/auth"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.Abort()
			return
		}

		headerParts := strings.Split(authHeader, " ")
		if len(headerParts) != 2 || headerParts[0] != "Bearer" {
			c.Abort()
			return
		}

		if len(headerParts[1]) == 0 {
			c.Abort()
			return
		}
		accessToken := headerParts[1]

		claims := &auth.Claims{}

		log.Println("AUTH")

		token, err := jwt.ParseWithClaims(accessToken, claims, func(token *jwt.Token) (interface{}, error) {
			return auth.JwtKey, nil
		})

		if err != nil || !token.Valid {
			if errors.Is(err, jwt.ErrSignatureInvalid) {
				c.JSON(http.StatusUnauthorized, gin.H{"message": "invalid token"})
				c.Abort()
				return
			}

			if ve, ok := err.(*jwt.ValidationError); ok && ve.Errors == jwt.ValidationErrorExpired {
				c.JSON(http.StatusUnauthorized, gin.H{"message": "token expired"})
				c.Abort()
				return
			}

			log.Println(err)

			c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
			c.Abort()
			return
		}

		c.Next()
	}
}
