package auth

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rakhmatullahyoga/mini-aspire/commons"
)

func ValidateJWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqToken := r.Header.Get("Authorization")
		splitToken := strings.Split(reqToken, "Bearer ")
		if len(splitToken) < 2 {
			res := commons.BuildErrorResponse("unauthorized request")
			w.WriteHeader(http.StatusUnauthorized)
			_ = json.NewEncoder(w).Encode(res)
			return
		}
		reqToken = splitToken[1]
		token, err := jwt.Parse(reqToken, func(t *jwt.Token) (interface{}, error) {
			return []byte(commons.JwtKey), nil
		})

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				res := commons.BuildErrorResponse("unauthorized request")
				w.WriteHeader(http.StatusUnauthorized)
				_ = json.NewEncoder(w).Encode(res)
				return
			}

			res := commons.BuildErrorResponse("bad request")
			w.WriteHeader(http.StatusBadRequest)
			_ = json.NewEncoder(w).Encode(res)
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			ctx := r.Context()
			ctx = context.WithValue(ctx, commons.ClaimsKeyUserID, claims[string(commons.ClaimsKeyUserID)])
			ctx = context.WithValue(ctx, commons.ClaimsKeyIsAdmin, claims[string(commons.ClaimsKeyIsAdmin)])
			next.ServeHTTP(w, r.WithContext(ctx))
		} else {
			res := commons.BuildErrorResponse("unauthorized request")
			w.WriteHeader(http.StatusUnauthorized)
			_ = json.NewEncoder(w).Encode(res)
			return
		}
	})
}

func EnsureAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		admin := ctx.Value(commons.ClaimsKeyIsAdmin).(bool)
		if !admin {
			res := commons.BuildErrorResponse("forbidden access")
			w.WriteHeader(http.StatusForbidden)
			_ = json.NewEncoder(w).Encode(res)
			return
		}

		next.ServeHTTP(w, r)
	})
}
