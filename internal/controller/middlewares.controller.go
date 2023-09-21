package controller

import (
	"net/http"

	"Uploader/internal/jwt-handler"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func jwtMiddleware(c *gin.Context, jwtService jwt_handler.Jwt) {
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
		c.Abort()

		return
	}

	token, err := jwt.ParseWithClaims(tokenString, &jwt_handler.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtService.SecretKey, nil
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

	c.Set("userID", claims.UserID)
}
