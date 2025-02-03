package user

import (
	"errors"
	"regexp"
)

var usernameValidator = regexp.MustCompile(`^[a-zA-Z0-9_ -]{1,30}$`)
var emailValidator = regexp.MustCompile(`^[^@]+@[^@]+\.[^@]{2,}$`)
var passwordValidator = regexp.MustCompile(`^.{8,64}$`)
var nameValidator = regexp.MustCompile(`^[a-zA-Z -]{1,80}$`)

type UserBase struct {
	Username  string `json:"username"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type UserDB struct {
	ID           int
	Email        string
	PasswordHash string
	UserBase
}

type GetUserResponse struct {
	ID int `json:"id"`
	UserBase
}

type GetCurrentUserResponse struct {
	Email string `json:"email"`
	GetUserResponse
}

type CreateUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	UserBase
}

type Credentials struct {
	Identifier string `json:"identifier"`
	Password   string `json:"password"`
}

type LoginResponse struct {
	Token    string                 `json:"token"`
	UserInfo GetCurrentUserResponse `json:"user_info"`
}

/*
 * Validate the CreateUserRequest
 * @param user: the CreateUserRequest to validate
 * @return error: an error if the CreateUserRequest is invalid
 */
func validateCreateUserRequest(user CreateUserRequest) (error) {
	if ok := usernameValidator.MatchString(user.Username); !ok {
		return errors.New("username")
	}
	if ok := emailValidator.MatchString(user.Email); !ok {
		return errors.New("email")
	}
	if ok := passwordValidator.MatchString(user.Password); !ok {
		return errors.New("password")
	}
	if ok := nameValidator.MatchString(user.FirstName); !ok {
		return errors.New("first_name")
	}
	if ok := nameValidator.MatchString(user.LastName); !ok {
		return errors.New("last_name")
	}

	return nil
}
