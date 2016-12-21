package repositories

import (
	"net/http"
	"github.com/auth-web-tokens/repositories/jwtback"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/dgrijalva/jwt-go"
	"fmt"
	"github.com/auth-web-tokens/models"
	"github.com/gorilla/context"
)

func RequireTokenAuthentication(rw http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
	authBackend := jwtback.InitAuthenticationBackend()

	token, err := request.ParseFromRequest(req, request.OAuth2Extractor, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		} else {
			return authBackend.PubicKey, nil
		}
	})

	if err == nil && token.Valid && !authBackend.IsInBlacklist(token.Raw) {
		uuid := authBackend.GetUserId(token)
		user := &models.User{}
		if err := user.GetById(uuid); err != nil {
			rw.WriteHeader(http.StatusUnauthorized)
			return
		}
		context.Set(req, "currentUser", user)
		next(rw, req)
	} else {
		rw.WriteHeader(http.StatusUnauthorized)
	}
}