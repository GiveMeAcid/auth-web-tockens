package repositories

import (
	"net/http"
	"github.com/auth-web-tokens/repositories/jwtback"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/dgrijalva/jwt-go"
	"fmt"
)

func RequireTokenAuthentication(rw http.ResponseWriter, req *http.Request, next http.Handler) {
	authBackend := jwtback.InitAuthenticationBackend()

	token, err := request.ParseFromRequest(req, request.OAuth2Extractor, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Fprintln("Unexpected signing method: %v", token.Header["alg"])
		} else {
			return authBackend.PubicKey, nil
		}
	})

	if err == nil && token.Valid && !authBackend.IsInBlacklist(req.Header.Get("Authorization")) {
		next(rw, req)
	} else {
		rw.WriteHeader(http.StatusUnauthorized)
	}
}