package auth

import (
	"errors"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecret []byte
var tokenLifetime time.Duration

/*
 * Init initializes the auth package with the jwt secret and token lifetime
 * @param tokenLifetime: the lifetime of the jwt token
 */
func Init(_tokenLifetime time.Duration, _jwtSecret []byte) {
	tokenLifetime = _tokenLifetime
	jwtSecret = _jwtSecret
}

/*
 * AuthorizeUser checks if the token in the http request authorization header is valid and returns the user id corresponding to the token
 * @param r: the http request
 * @return int: the user id
 * @return error: an error if the user is not authorized
 */
func AuthorizeUser(r *http.Request) (int, error) {
	tokenString := r.Header.Get("Authorization")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
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

/*
 * HashPassword hashes the password using bcrypt
 * @param password: the password to hash
 * @return string: the hashed password
 * @return error: an error if the password could not be hashed
 */
func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}

/*
 * ValidatePassword checks if the password matches the hash
 * @param password: the password to check
 * @param hash: the hash to check against
 * @return error: an error if the password does not match the hash
 */
func ValidatePassword(password, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

/*
 * GenerateToken generates a jwt token with the user id
 * @param id: the user id
 * @return string: the jwt token
 * @return error: an error if the token could not be generated
 */
func GenerateToken(id int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": id,
		"exp":     time.Now().Add(tokenLifetime).Unix(),
	})
	return token.SignedString(jwtSecret)
}
