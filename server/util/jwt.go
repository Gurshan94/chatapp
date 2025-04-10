package util

import (
    "errors"
	"github.com/golang-jwt/jwt/v4"
)

const (
	secretKey = "secret"
)

type MyJWTClaims struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email string `json:"email"`

	jwt.RegisteredClaims
}



func ValidateJWT(tokenString string) (*MyJWTClaims, error) {
    token, err := jwt.ParseWithClaims(tokenString, &MyJWTClaims{}, func(token *jwt.Token) (interface{}, error) {
        // Ensure the token method conforms to "SigningMethodHMAC"
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, errors.New("unexpected signing method")
        }
        return []byte(secretKey), nil
    })

    if err != nil {
        return nil, err
    }

    // Extract the claims
    claims, ok := token.Claims.(*MyJWTClaims)
    if ok && token.Valid {
        return claims, nil
    }

    return nil, errors.New("invalid token")
}