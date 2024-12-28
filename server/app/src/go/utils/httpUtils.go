package utils

import (
	"errors"
	"net/http"
	"encoding/json"

	"software-slayer/auth"
)

func Decode[T any](w http.ResponseWriter, r *http.Request, v *T) error {
	err := json.NewDecoder(r.Body).Decode(v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	return err
}

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