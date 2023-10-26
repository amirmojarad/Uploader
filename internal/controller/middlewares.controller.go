package controller

import (
	"Uploader/consts"
	"net/http"
	"strings"

	"Uploader/internal/jwt-handler"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type Middleware struct {
	jwt *jwt_handler.Jwt
}

func NewMiddleware(jwt *jwt_handler.Jwt) *Middleware {
	return &Middleware{
		jwt: jwt,
	}
}

func (m Middleware) jwtMiddleware(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
		c.Abort()

		return
	}

	splitToken := strings.Split(tokenString, " ")

	token, err := jwt.ParseWithClaims(splitToken[1], &jwt_handler.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return m.jwt.SecretKey, nil
	})

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		c.Abort()

		return
	}

	claims, ok := token.Claims.(*jwt_handler.Claims)
	if !ok || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		c.Abort()

		return
	}

	c.Set(consts.REQUEST_USER_ID, claims.UserID)
}
