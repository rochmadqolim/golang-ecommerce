package middleware

import (
	"net/http"

	"github.com/golang-jwt/jwt/v4"
	"github.com/rochmadqolim/golang-ecommerce/config"
	"github.com/rochmadqolim/golang-ecommerce/responses"
)

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie("token")
		if err != nil {
			response := map[string]string{"error": "Authentication required"}
			responses.ResponseJSON(w, http.StatusUnauthorized, response)
			return
		}

		// Get JWT token from cookie
		tokenString := c.Value

		// Parsing JWT token
		claims := &config.JWTClaim{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
			return config.JWTKey, nil
		})

		if err != nil {
			if ve, ok := err.(*jwt.ValidationError); ok {
				if ve.Errors&jwt.ValidationErrorSignatureInvalid != 0 {
					response := map[string]string{"error": "Invalid token"}
					responses.ResponseJSON(w, http.StatusUnauthorized, response)
					return
				} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
					response := map[string]string{"error": "Token expired"}
					responses.ResponseJSON(w, http.StatusUnauthorized, response)
					return
				}
			}

			response := map[string]string{"error": "Authentication failed"}
			responses.ResponseJSON(w, http.StatusUnauthorized, response)
			return
		}

		if token.Valid {
			// Token is valid, continue to the next handler
			next.ServeHTTP(w, r)
		} else {
			response := map[string]string{"error": "Authentication failed"}
			responses.ResponseJSON(w, http.StatusUnauthorized, response)
			return
		}
	})
}
