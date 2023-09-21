package jwt_handler

import (
	"time"

	"Uploader/internal/errorext"
	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	UserID uint `json:"user_id"`
	jwt.StandardClaims
}

type Jwt struct {
	SecretKey      []byte
	ExpireDuration int
}

func NewJwt(secretKey string, expireDuration int) *Jwt {
	return &Jwt{
		SecretKey:      []byte(secretKey),
		ExpireDuration: expireDuration,
	}
}

func (j Jwt) GenerateJWT(userID uint) (string, error) {
	claims := &Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * time.Duration(j.ExpireDuration)).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(j.SecretKey)

	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func (j Jwt) ValidateJWT(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SecretKey, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, errorext.NewTokenError("invalid token")
	}

	return claims, nil
}
