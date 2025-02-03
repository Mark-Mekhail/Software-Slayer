package auth_test

import (
	"testing"
	"time"

	"software-slayer/auth"
)

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
		tokenService := auth.NewTokenService(time.Duration(duration), []byte(jwtSecret))

		token, err := tokenService.GenerateToken(userId)
		if err != nil {
			t.Errorf("GenerateToken() returned error: %v", err)
		}

		id, err := tokenService.AuthorizeUser(token)
		if err != nil {
			t.Errorf("AuthorizeUser() returned error: %v", err)
		}

		if id != userId {
			t.Errorf("AuthorizeUser() returned wrong user id: %d", id)
		}
	})
}

func TestAuthorizeUserExpiredToken(t *testing.T) {
	tokenService := auth.NewTokenService(time.Second, []byte("secret"))

	token, err := tokenService.GenerateToken(1)
	if err != nil {
		t.Errorf("GenerateToken() returned error: %v", err)
	}

	_, err = tokenService.AuthorizeUser(token)
	if err != nil {
		t.Errorf("AuthorizeUser() returned error: %v", err)
	}

	time.Sleep(time.Second * 2)

	_, err = tokenService.AuthorizeUser(token)
	if err == nil {
		t.Errorf("AuthorizeUser() did not return error for expired token")
	}
}