package auth

import (
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func CreateJWT(id string, email string) (string, error) {
	expiration := time.Now().Add(168 * time.Hour)
	claims := &Claims{
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        id,
			ExpiresAt: jwt.NewNumericDate(expiration),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secretKey)
	return tokenString, err
}

func GetJWT(tokenStr string) (int, bool) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (any, error) {
		return secretKey, nil
	})
	if err != nil {
		return 0, false
	}
	id, err := strconv.Atoi(claims.ID)
	if err != nil {
		return 0, false
	}
	return id, token.Valid
}
