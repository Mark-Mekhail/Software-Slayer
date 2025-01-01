package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"regexp"

	"software-slayer/auth"
)

/*
 * Decode decodes the request body into the given struct.
 * @param w: the response writer
 * @param r: the request
 * @param v: the struct to decode the request body into
 * @return error: an error if the decoding fails
 */
func Decode[T any](w http.ResponseWriter, r *http.Request, v *T) error {
	err := json.NewDecoder(r.Body).Decode(&v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	return err
}

/*
 * ValidateUser validates the user making the request.
 * @param w: the response writer
 * @param r: the request
 * @param expectedUserId: the expected user id
 * @return error: an error if the user is not authorized
 */
func ValidateUser(w http.ResponseWriter, r *http.Request, expectedUserId int) error {
	user_id, err := auth.AuthorizeUser(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
	} else if user_id != expectedUserId {
		err = errors.New("unauthorized")
		http.Error(w, err.Error(), http.StatusUnauthorized)
	}
	return err
}
