package auth

import (
	"errors"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
)

type TokenService struct {
	jwtSecret     []byte
	tokenLifetime time.Duration
}

func NewTokenService(tokenLifetime time.Duration, jwtSecret []byte) *TokenService {
	return &TokenService{tokenLifetime: tokenLifetime, jwtSecret: jwtSecret}
}

func (s *TokenService) AuthorizeUser(r *http.Request) (int, error) {
	tokenString := r.Header.Get("Authorization")

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

func (s *TokenService) GenerateToken(id int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": id,
		"exp":     time.Now().Add(s.tokenLifetime).Unix(),
	})
	return token.SignedString(s.jwtSecret)
}
