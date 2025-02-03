package user

import (
	"software-slayer/db"
)

type UserService interface {
	CreateUser(user *CreateUserRequest, passwordHash string) error
	GetUsers() ([]GetUserResponse, error)
	GetUserByIdentifier(identifier string) (UserDB, error)
	GetUserById(id int) (UserDB, error)
}

type UserServiceImpl struct {
	db *db.Database
}

func NewUserService(db *db.Database) *UserServiceImpl {
	return &UserServiceImpl{db: db}
}

func (s *UserServiceImpl) CreateUser(user *CreateUserRequest, passwordHash string) error {
	_, err := s.db.Exec("INSERT INTO users (email, username, password_hash, first_name, last_name) VALUES (?, ?, ?, ?, ?)", user.Email, user.Username, passwordHash, user.FirstName, user.LastName)
	return err
}

func (s *UserServiceImpl) GetUsers() ([]GetUserResponse, error) {
	rows, err := s.db.Query("SELECT id, username, first_name, last_name FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := make([]GetUserResponse, 0)
	for rows.Next() {
		var user GetUserResponse
		err := rows.Scan(&user.ID, &user.Username, &user.FirstName, &user.LastName)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (s *UserServiceImpl) GetUserByIdentifier(identifier string) (UserDB, error) {
	var user UserDB
	err := s.db.QueryRow("SELECT id, username, email, password_hash, first_name, last_name FROM users WHERE email = ? OR username = ?", identifier, identifier).Scan(&user.ID, &user.Username, &user.Email, &user.PasswordHash, &user.FirstName, &user.LastName)
	return user, err
}

func (s *UserServiceImpl) GetUserById(id int) (UserDB, error) {
	var user UserDB
	err := s.db.QueryRow("SELECT id, username, email, password_hash, first_name, last_name FROM users WHERE id = ?", id).Scan(&user.ID, &user.Username, &user.Email, &user.PasswordHash, &user.FirstName, &user.LastName)
	return user, err
}
