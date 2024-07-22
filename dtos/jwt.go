package dtos

import (
	"github.com/golang-jwt/jwt/v4"
)

var JWT_KEY = []byte("nikonug")

type JWTClaim struct {
	Email string
	jwt.RegisteredClaims
}
