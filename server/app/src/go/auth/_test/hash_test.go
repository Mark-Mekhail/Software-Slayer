package auth_test

import (
	"testing"

	"software-slayer/auth"
)

func FuzzHashPassword(f *testing.F) {
	passwordSeeds := []string{ "password", "pass", "p", "verylongpasswordthatjustkeepsgettinglongerandlonger", "78934218057923", "()_*&(*%^&$#(*#)*%$#><?>{L{}})", "" }
	for _, password := range passwordSeeds {
		f.Add(password)
	}

	f.Fuzz(func(t *testing.T, password string) {
		hashedPassword, err := auth.HashPassword(password)
		if err != nil {
			t.Errorf("HashPassword(%s) returned error: %v", password, err)
		}

		if err := auth.ValidatePassword(password, hashedPassword); err != nil {
			t.Errorf("ValidatePassword(%s, %s) returned false", password, hashedPassword)
		}
	})
}