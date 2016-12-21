package services

import (
	"github.com/auth-web-tokens/models/requests"
	"github.com/auth-web-tokens/models"
	"github.com/auth-web-tokens/repositories/jwtback"
	"net/http"
	"encoding/json"
	"github.com/auth-web-tokens/app/parametrs"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/dgrijalva/jwt-go"
)

func Login(requestUser requests.User, dbUser models.User) (int, string){
	authBackend :=  jwtback.InitAuthenticationBackend()

	if authBackend.Authenticate(requestUser, dbUser) {
		token, err := authBackend.GenerateToken(dbUser.UUID)
		if err != nil {
			return http.StatusInternalServerError, ""
		} else {
			return http.StatusOK, token
		}
	}
	return http.StatusUnauthorized, ""
}

func RefreshToken(requestUser requests.User, user *models.User) []byte {
	authBackend := jwtback.InitAuthenticationBackend()
	token, err := authBackend.GenerateToken(user.UUID)
	if err != nil {
		panic(err)
	}

	response, err := json.Marshal(parametrs.TokenAuthentication{Token: token})
	if err != nil {
		panic(err)
	}
	return response
}

func Logout(req http.Request) error {
	authBackend := jwtback.InitAuthenticationBackend()
	tokenRequest, err := request.ParseFromRequest(req, request.OAuth2Extractor, func(token *jwt.Token) (interface{}, error){
		return authBackend.PubicKey, nil
	})
	if err != nil {
		return err
	}
	tokenString := req.Header.Get("Authorization")
	return authBackend.Logout(tokenString, tokenRequest)
}