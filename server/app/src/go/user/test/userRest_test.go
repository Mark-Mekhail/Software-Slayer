package user_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"software-slayer/user"
	"software-slayer/auth"
)

type MockUserService struct{}

func (m *MockUserService) CreateUser(user *user.CreateUserRequest, passwordHash string) error {
	if user.Email == "invalid" {
		return errors.New("invalid email")
	}
	return nil
}

func (m *MockUserService) GetUsers() ([]user.GetUserResponse, error) {
	return []user.GetUserResponse{
		{
			ID:
			1,
			UserBase: user.UserBase{
				Username:  "testuser",
				FirstName: "John",
				LastName:  "Doe",
			},
		},
	}, nil
}

func (m *MockUserService) GetUserByIdentifier(identifier string) (user.UserDB, error) {
	hashedPassword, _ := auth.HashPassword("password123")

	if identifier == "test@example.com" {
		return user.UserDB{
			ID: 1,
			Email: "test@example.com",
			PasswordHash: hashedPassword,
		}, nil
	}
	return user.UserDB{}, errors.New("user not found")
}

func (m *MockUserService) GetUserById(id int) (user.UserDB, error) {
	if id == 1 {
		return user.UserDB{
			ID: 1,
			Email: "test@example.com",
			PasswordHash: "hashedpassword",
		}, nil
	}
	return user.UserDB{}, errors.New("user not found")
}

type MockTokenService struct{}

func (m *MockTokenService) GenerateToken(userID int) (string, error) {
	return "mocked_token", nil
}

func (m *MockTokenService) AuthorizeUser(token string) (int, error) {
	if token == "valid_token" {
		return 1, nil
	}
	return 0, errors.New("invalid token")
}

func SetupServer() *httptest.Server {
	mockUserService := &MockUserService{}
	mockTokenService := &MockTokenService{}
	user.InitUserRest(mockUserService, mockTokenService)
	return httptest.NewServer(http.DefaultServeMux)
}

func TestCreateUserSuccess(t *testing.T) {
	ts := SetupServer()
	defer ts.Close()

	requestBody := user.CreateUserRequest{
		Email:    "test@example.com",
		Password: "password123",
		UserBase: user.UserBase{
			Username:  "testuser",
			FirstName: "John",
			LastName:  "Doe",
		},
	}
	body, _ := json.Marshal(requestBody)

	resp, err := http.Post(ts.URL+"/user", "application/json", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		t.Errorf("expected %d, got %d", http.StatusCreated, resp.StatusCode)
	}
}

func TestHandleLoginSuccess(t *testing.T) {
	ts := SetupServer()
	defer ts.Close()

	requestBody := user.Credentials{
		Identifier: "test@example.com",
		Password: "password123",
	}
	body, _ := json.Marshal(requestBody)

	resp, err := http.Post(ts.URL+"/login", "application/json", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected %d, got %d", http.StatusOK, resp.StatusCode)
	}
}

func TestGetAllUsers(t *testing.T) {
	ts := SetupServer()
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/user")
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected %d, got %d", http.StatusOK, resp.StatusCode)
	}
}

func TestGetCurrentUser(t *testing.T) {
	ts := SetupServer()
	defer ts.Close()

	req, _ := http.NewRequest("GET", ts.URL+"/user", nil)
	req.Header.Set("Authorization", "Bearer valid_token")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected %d, got %d", http.StatusOK, resp.StatusCode)
	}
}
