package common

import (
	"crypto/rand"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

type TokenProvider interface {
	GenRefreshToken() (string, error)
	GenAccessToken(userId int, expiry int) (string, error)
	ParseAccessToken(tokenString string) (*jwt.Token, error)
}

type DefaultTokenProvider struct {
	secretKey string
}

func NewDefaultTokenProvider(secretKey string) *DefaultTokenProvider {
	return &DefaultTokenProvider{secretKey: secretKey}
}

func (tkp *DefaultTokenProvider) GenRefreshToken() (string, error) {
	b := make([]byte, 8)
	rand.Read(b)
	return fmt.Sprintf("%x", b), nil
}

func (tkp *DefaultTokenProvider) GenAccessToken(userId int, expiry int) (string, error) {
	var err error
	atClaims := jwt.MapClaims{}
	atClaims["user_id"] = userId
	atClaims["exp"] = time.Now().Add(time.Minute * time.Duration(expiry)).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(tkp.secretKey))

	if err != nil {
		return "", err
	}

	return token, nil
}

func (tkp *DefaultTokenProvider) ParseAccessToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(tkp.secretKey), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}
