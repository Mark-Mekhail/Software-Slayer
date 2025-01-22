package user_test

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"

	"software-slayer/user"
)

type MockDatabase struct {
	mock.Mock
}

func (m *MockDatabase) Close() error {
	arguments := m.Called()
	return arguments.Error(0)
}

func (m *MockDatabase) Exec(query string, args ...any) (sql.Result, error) {
	arguments := m.Called(query, args)
	return nil, arguments.Error(1)
}

func (m *MockDatabase) QueryRow(query string, args ...any) *sql.Row {
	arguments := m.Called(query, args)
	return arguments.Get(0).(*sql.Row)
}

func (m *MockDatabase) Query(query string, args ...any) (*sql.Rows, error) {
	arguments := m.Called(query, args)
	return arguments.Get(0).(*sql.Rows), arguments.Error(1)
}

func TestCreateUser(t *testing.T) {
	mockDB := new(MockDatabase)
	s := user.NewUserService(mockDB)

	mockDB.On("Exec", mock.Anything, mock.Anything).Return(nil, nil)

	user := &user.CreateUserRequest{
		Email: "user@gmail.com",
		UserBase: user.UserBase{
			Username:  "user",
			FirstName: "John",
			LastName:  "Doe",
		},
	}
	passwordHash := "passwordHash"

	err := s.CreateUser(user, passwordHash)
	if err != nil {
		t.Error("Expected nil, got", err)
	}

	mockDB.AssertExpectations(t)
}

func TestCreateUserError(t *testing.T) {
	mockDB := new(MockDatabase)
	s := user.NewUserService(mockDB)

	mockDB.On("Exec", mock.Anything, mock.Anything).Return(nil, errors.New("error"))

	user := &user.CreateUserRequest{
		Email: "user@gmail.com",
		UserBase: user.UserBase{
			Username:  "user",
			FirstName: "John",
			LastName:  "Doe",
		},
	}
	passwordHash := "passwordHash"

	err := s.CreateUser(user, passwordHash)
	if err == nil {
		t.Error("Expected error, got nil")
	}

	mockDB.AssertExpectations(t)
}