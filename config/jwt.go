package config

import "github.com/golang-jwt/jwt/v5"

var JWT_KEY = []byte("ashdjqy9283409bsdklkg8hda01")

type JWTClaim struct {
	Email string
	jwt.RegisteredClaims
}