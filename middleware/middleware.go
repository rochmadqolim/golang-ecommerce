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
			response := map[string]string{"message": "Unauthorized"}
			responses.ResponseJSON(w, http.StatusUnauthorized, response)
			return
		}

		// Mengambil token value
		tokenString := c.Value

		claims := &config.JWTClaim{}
		// Parsing token JWT
		token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
			return config.JWTKey, nil
		})

		if err != nil {
			validate, _ := err.(*jwt.ValidationError)
			switch validate.Errors {
			case jwt.ValidationErrorSignatureInvalid:
				// Token tidak valid
				response := map[string]string{"message": "Unauthorized"}
				responses.ResponseJSON(w, http.StatusUnauthorized, response)
				return
			case jwt.ValidationErrorExpired:
				// Token telah kadaluarsa
				response := map[string]string{"message": "Unauthorized, Token expired!"}
				responses.ResponseJSON(w, http.StatusUnauthorized, response)
				return
			default:
				response := map[string]string{"message": "Unauthorized"}
				responses.ResponseJSON(w, http.StatusUnauthorized, response)
				return
			}
		}

		if !token.Valid {
			response := map[string]string{"message": "Unauthorized"}
			responses.ResponseJSON(w, http.StatusUnauthorized, response)
			return
		}

		next.ServeHTTP(w, r)
	})
}
