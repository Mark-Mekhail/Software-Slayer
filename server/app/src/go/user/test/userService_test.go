package user_test

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"

	"software-slayer/db"
	"software-slayer/user"
)

func setup(t *testing.T) (sqlmock.Sqlmock, *user.UserService) {
	// Create a mock sql.DB object using sqlmock
	database, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	s := user.NewUserService(*db.NewTestDB(database))

	return mock, s
}

func TestCreateUser(t *testing.T) {
	dbMock, s := setup(t)

	user := &user.CreateUserRequest{
		Email: "user@gmail.com",
		UserBase: user.UserBase{
			Username:  "user",
			FirstName: "John",
			LastName:  "Doe",
		},
	}
	passwordHash := "passwordHash"

	dbMock.ExpectExec("INSERT INTO users").WithArgs(user.Email, user.Username, passwordHash, user.FirstName, user.LastName).WillReturnResult(sqlmock.NewResult(1, 1))

	err := s.CreateUser(user, passwordHash)
	if err != nil {
		t.Error("Expected nil, got ", err)
	}
}

func TestCreateUserError(t *testing.T) {
	dbMock, s := setup(t)

	user := &user.CreateUserRequest{
		Email: "user@gmail.com",
		UserBase: user.UserBase{
			Username:  "user",
			FirstName: "John",
			LastName:  "Doe",
		},
	}
	passwordHash := "passwordHash"

	dbMock.ExpectExec("INSERT INTO users").WithArgs(user.Email, user.Username, user.FirstName, user.LastName).WillReturnError(errors.New("error"))

	err := s.CreateUser(user, passwordHash)
	if err == nil {
		t.Error("Expected error, got nil")
	}
}

func TestGetUsers(t *testing.T) {
	dbMock, s := setup(t)

	users := []user.GetUserResponse{
		{
			ID: 1, 
			UserBase: user.UserBase{
				Username: "user1", 
				FirstName: "John", 
				LastName: "Doe",
			},
		},
		{
			ID: 2,
			UserBase: user.UserBase{
				Username: "user2",
				FirstName: "Jane",
				LastName: "Doe",
			},
		},
	}

	rows := sqlmock.NewRows([]string{"id", "username", "first_name", "last_name"})
	for _, user := range users {
		rows.AddRow(user.ID, user.Username, user.FirstName, user.LastName)
	}

	dbMock.ExpectQuery("SELECT id, username, first_name, last_name FROM users").WillReturnRows(rows)

	res, err := s.GetUsers()
	if err != nil {
		t.Error("Expected nil, got ", err)
	}

	if len(res) != len(users) {
		t.Error("Expected 2 users, got ", len(users))
	}

	for i, user := range users {
		if user.ID != res[i].ID || user.Username != res[i].Username || user.FirstName != res[i].FirstName || user.LastName != res[i].LastName {
			t.Error("Expected ", user, ", got ", res[i])
		}
	}
}

func TestGetNoUsers(t *testing.T) {
	dbMock, s := setup(t)

	rows := sqlmock.NewRows([]string{"id", "username", "first_name", "last_name"})

	dbMock.ExpectQuery("SELECT id, username, first_name, last_name FROM users").WillReturnRows(rows)

	res, err := s.GetUsers()
	if err != nil {
		t.Error("Expected nil, got ", err)
	}

	if len(res) != 0 {
		t.Error("Expected 0 users, got ", len(res))
	}
}

func TestGetUsersError(t *testing.T) {
	dbMock, s := setup(t)

	dbMock.ExpectQuery("SELECT id, username, first_name, last_name FROM users").WillReturnError(errors.New("error"))

	_, err := s.GetUsers()
	if err == nil {
		t.Error("Expected error, got nil")
	}
}