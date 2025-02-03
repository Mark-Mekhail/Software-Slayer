package utils

import (
	"encoding/json"
	"net/http"
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
