package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

type TokenService interface {
	AuthorizeUser(tokenString string) (int, error)
	GenerateToken(id int) (string, error)
}

type TokenServiceImpl struct {
	jwtSecret     []byte
	tokenLifetime time.Duration
}

func NewTokenService(tokenLifetime time.Duration, jwtSecret []byte) *TokenServiceImpl {
	return &TokenServiceImpl{tokenLifetime: tokenLifetime, jwtSecret: jwtSecret}
}

func (s *TokenServiceImpl) AuthorizeUser(tokenString string) (int, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return s.jwtSecret, nil
	})
	if err != nil || !token.Valid {
		return -1, errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return -1, errors.New("invalid claims")
	}

	userId, ok := claims["user_id"]
	if !ok {
		return -1, errors.New("invalid user id")
	}

	return int(userId.(float64)), nil
}

func (s *TokenServiceImpl) GenerateToken(id int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": id,
		"exp":     time.Now().Add(s.tokenLifetime).Unix(),
	})
	return token.SignedString(s.jwtSecret)
}
