package middleware

import (
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

/*
func ValidateJWT(next func(w http.ResponseWriter, r *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header["Token"] != nil {
			token, err := jwt.Parse(r.Header["Token"][0], func(t *jwt.Token) (interface{}, error) {
				_, ok := t.Method.(*jwt.SigningMethodHMAC)
				if !ok {
					w.WriteHeader(http.StatusUnauthorized)
					w.Write([]byte("unauthorized"))
				}
				return SECRET, nil

			})
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("unauthorized : " + err.Error()))
			}
			if token.Valid {
				next(w, r)
			}
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("unauthorized"))
		}
	})
}
*/

func ValidateJWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check for the "Authorization" header
		authHeader := r.Header.Get("Authorization")
		if authHeader != "" {
			// Split the header value into "Bearer" and the token
			parts := strings.Split(authHeader, " ")
			if len(parts) != 2 || parts[0] != "Bearer" {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("Unauthorized: Invalid token format"))
				return
			}

			tokenString := parts[1]

			token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
				_, ok := t.Method.(*jwt.SigningMethodHMAC)
				if !ok {
					w.WriteHeader(http.StatusUnauthorized)
					w.Write([]byte("Unauthorized: Invalid token signing method"))
				}
				return SECRET, nil
			})

			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("Unauthorized: " + err.Error()))
				return
			}

			if token.Valid {
				next.ServeHTTP(w, r)
				return
			}
		}

		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Unauthorized: Missing or invalid token"))
	})
}
