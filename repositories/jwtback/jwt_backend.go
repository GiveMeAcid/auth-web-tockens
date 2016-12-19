package jwtback

import (
	"crypto/rsa"
	"github.com/dgrijalva/jwt-go"
	"time"
	"github.com/satori/go.uuid"
	"github.com/auth-web-tokens/settings"
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

func getPrivateKey() *rsa.PrivateKey {
	return
}

func getPublicKey() *rsa.PublicKey {
	return
}