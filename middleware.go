package main

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
)

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			respondWithError("Access not allowed", errors.New("Invalid jwt token"), w, http.StatusForbidden)
			return
		}

		// TODO: Make this a bit more robust, parsing-wise
		authHeaderParts := strings.Split(authHeader, " ")
		if len(authHeaderParts) != 2 || strings.ToLower(authHeaderParts[0]) != "bearer" {
			respondWithError("Access not allowed", errors.New("Authorization header format must be Bearer {token}"), w, http.StatusForbidden)
			return
		}

		jwtToken := authHeaderParts[1]

		token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
			// Valid alg is what we expect
			if token.Method != jwt.SigningMethodHS256 {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(SecretKey), nil
		})
		if err != nil {
			respondWithError("Access not allowed", err, w, http.StatusForbidden)
			return
		}

		if !token.Valid {
			respondWithError("Access not allowed", errors.New("Invalid jwt token"), w, http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}
