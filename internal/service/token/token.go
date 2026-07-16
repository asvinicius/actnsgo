package token

import (
	"errors"
	"time"

	"github.com/asvinicius/actnsgo/internal/model"
	"github.com/golang-jwt/jwt/v5"
)

type TokenService struct {
	secret     []byte
	expiration time.Duration
}

func NewTokenService(secret, expiration string) (*TokenService, error) {
	duration, err := time.ParseDuration(expiration)

	if err != nil {
		return nil, err
	}

	return &TokenService{
		secret:     []byte(secret),
		expiration: duration,
	}, nil
}

func (s *TokenService) GenerateSuperToken(super *model.UserSuper) (string, error) {
	exp := time.Now().Add(s.expiration)

	claims := jwt.MapClaims{
		"sub":   super.SuperID,
		"login": super.SuperLogin,
		"type":  "super",
		"exp":   exp.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(s.secret)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (s *TokenService) ValidateToken(tokenString string) (jwt.MapClaims, error) {
	claims := jwt.MapClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (any, error) {
		if token.Method != jwt.SigningMethodHS256 {
			return nil, errors.New("método de assinatura inválido")
		}
		return s.secret, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("token inválido")
	}

	return claims, nil
}
