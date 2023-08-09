package config

import (
	"os"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte(os.Getenv("JWT_KEY"))

type JWTClaim struct {
	Email string
	jwt.RegisteredClaims
}