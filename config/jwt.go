package config

import (
	"os"

	"github.com/golang-jwt/jwt/v4"
)

// JWTKey to verification JWT token
var JWTKey = []byte(os.Getenv("JWT_KEY"))

type JWTClaim struct {
	Email string
	jwt.RegisteredClaims
}