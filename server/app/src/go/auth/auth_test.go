package auth

import (
	"net/http"
	"testing"
	"time"
)

func FuzzHashPassword(f *testing.F) {
	secretSeeds := []string{ "secret", "seed", "s", "verylongsecretthatjustkeepsgettinglongerandlonger", "78934218057923", "()_*&(*%^&$#(*#)*%$#><?>{L{}})" }
	passwordSeeds := []string{ "password", "pass", "p", "verylongpasswordthatjustkeepsgettinglongerandlonger", "78934218057923", "()_*&(*%^&$#(*#)*%$#><?>{L{}})", "" }

	for _, secret := range secretSeeds {
		for _, password := range passwordSeeds {
			f.Add(secret, password)
		}
	}

	f.Fuzz(func(t *testing.T, jwtSecret string, password string) {
		Init(time.Minute, []byte(jwtSecret))

		hashedPassword, err := HashPassword(password)
		if err != nil {
			t.Errorf("HashPassword(%s) returned error: %v", password, err)
		}

		if err := ValidatePassword(password, hashedPassword); err != nil {
			t.Errorf("ValidatePassword(%s, %s) returned false", password, hashedPassword)
		}
	})
}

func FuzzAuthorizeUser(f *testing.F) {
	durationSeeds := []int64{ int64(time.Hour), int64(time.Minute), int64(3 * time.Second) }
	secretSeeds := []string{ "s", "secret", "superlongsecretthatkeepsgettinglongerandlonger", "&*(%$#)@)(#*%)}:>?][]|" }
	userIdSeeds := []int{ 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 123, 985, 9999999999, 843279 }

	for _, duration := range durationSeeds {
		for _, secret := range secretSeeds {
			for _, userId := range userIdSeeds {
				f.Add(duration, secret, userId)
			}
		}
	}

	f.Fuzz(func(t *testing.T, duration int64, jwtSecret string, userId int) {
		Init(time.Duration(duration), []byte(jwtSecret))

		token, err := GenerateToken(userId)
		if err != nil {
			t.Errorf("GenerateToken() returned error: %v", err)
		}
		
		r := &http.Request{
			Header: map[string][]string{
				"Authorization": {token},
			},
		}

		id, err := AuthorizeUser(r)
		if err != nil {
			t.Errorf("AuthorizeUser() returned error: %v", err)
		}

		if id != userId {
			t.Errorf("AuthorizeUser() returned wrong user id: %d", id)
		}
	})
}

func TestAuthorizeUserExpiredToken(t *testing.T) {
	Init(time.Second, []byte("secret"))

	token, err := GenerateToken(1)
	if err != nil {
		t.Errorf("GenerateToken() returned error: %v", err)
	}

	r := &http.Request{
		Header: map[string][]string{
			"Authorization": {token},
		},
	}

	_, err = AuthorizeUser(r)
	if err != nil {
		t.Errorf("AuthorizeUser() returned error: %v", err)
	}

	time.Sleep(time.Second * 2)

	_, err = AuthorizeUser(r)
	if err == nil {
		t.Errorf("AuthorizeUser() did not return error for expired token")
	}
}