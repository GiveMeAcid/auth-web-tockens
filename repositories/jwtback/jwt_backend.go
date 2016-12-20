package jwtback

import (
	"crypto/rsa"
	"github.com/dgrijalva/jwt-go"
	"time"
	"github.com/satori/go.uuid"
	"github.com/auth-web-tokens/settings"
	"github.com/auth-web-tokens/models/requests"
	"github.com/auth-web-tokens/models"
	"golang.org/x/crypto/bcrypt"
)

type JWTAuthenticationBackend struct {
	privateKey *rsa.PrivateKey
	PubicKey   *rsa.PublicKey
}

const (
	tokenDuration = 72
	expireOffset = 3600
)

var authBackendInstance *JWTAuthenticationBackend = nil

func InitAuthenticationBackend() *JWTAuthenticationBackend {
	if authBackendInstance == nil {
		authBackendInstance = &JWTAuthenticationBackend{
			privateKey: getPrivateKey(),
			PubicKey: getPublicKey(),
		}
	}
	return authBackendInstance
}

func (backend *JWTAuthenticationBackend) GenerateToken(uuid uuid.UUID) (string, error) {
	token := jwt.New(jwt.SigningMethodRS512)
	token.Claims = &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * time.Duration(settings.Get().JWTExpirationDelta)).Unix(),
		IssuedAt: time.Now().Unix(),
		Subject: uuid.String(),
	}
	tokenString, err := token.SignedString(backend.privateKey)
	if err != nil {
		panic(err)
		return "", err
	}
	return tokenString, nil
}

func (backend *JWTAuthenticationBackend) Authenticate(user *requests.User, dbUser *models.User) bool {
	return user.Email ==  dbUser.Email && bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password)) == nil
}

func getPrivateKey() *rsa.PrivateKey {
	return
}

func getPublicKey() *rsa.PublicKey {
	return
}