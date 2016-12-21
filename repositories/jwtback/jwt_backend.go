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
	"github.com/auth-web-tokens/services/redis"
	"io/ioutil"
	"encoding/pem"
	"crypto/x509"
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

func (backend *JWTAuthenticationBackend) GetUserId(token *jwt.Token) uuid.UUID {
	claims := token.Claims.(*jwt.StandardClaims)
	return uuid.FromStringOrNil(claims.Subject)
}

func (backend *JWTAuthenticationBackend) Authenticate(user *requests.User, dbUser *models.User) bool {
	return user.Email == dbUser.Email && bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password)) == nil
}

func (backend *JWTAuthenticationBackend) getTokenRemainingValidity(timestamp interface{}) int {
	if validity, ok := timestamp.(float64); ok {
		tm := time.Unix(int64(validity), 0)
		remainer := tm.Sub(time.Now())
		if remainer > 0 {
			return int(remainer.Seconds() + expireOffset)
		}
	}
	return expireOffset
}

func (backend *JWTAuthenticationBackend) Logout(tokenString string, token *jwt.Token) error {
	claims := token.Claims.(jwt.MapClaims)
	return redis.GetInstance().SetValue(tokenString, tokenString, backend.getTokenRemainingValidity(claims["exp"]))
}

func (backend *JWTAuthenticationBackend) IsInBlacklist(token string) bool {
	redisToken, _ := redis.GetInstance().GetValue(token)

	if redisToken == nil {
		return false
	}

	return true
}

func getPrivateKey() *rsa.PrivateKey {
	pembytes, err := ioutil.ReadFile(settings.Get().PrivateKeyPath)
	if err != nil {
		panic(err)
	}
	data, _ := pem.Decode([]byte(pembytes))
	privateKeyImported, err := x509.ParsePKCS1PrivateKey(data.Bytes)

	if err != nil {
		panic(err)
	}

	return privateKeyImported
}

func getPublicKey() *rsa.PublicKey {
	pembytes, err := ioutil.ReadFile(settings.Get().PublicKeyPath)
	if err != nil {
		panic(err)
	}

	data, _ := pem.Decode([]byte(pembytes))
	publicKeyImported, err := x509.ParsePKIXPublicKey(data.Bytes)
	if err != nil {
		panic(err)
	}

	rsaPub, ok := publicKeyImported.(*rsa.PublicKey)

	if !ok {
		panic(err)
	}

	return rsaPub
}