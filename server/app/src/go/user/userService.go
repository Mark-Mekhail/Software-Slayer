package user

import (
	"software-slayer/db"
)

type UserService struct {
	db *db.Database
}

func NewUserService(db *db.Database) *UserService {
	return &UserService{db: db}
}

func (s *UserService) CreateUser(user *CreateUserRequest, passwordHash string) error {
	_, err := s.db.Exec("INSERT INTO users (email, username, password_hash, first_name, last_name) VALUES (?, ?, ?, ?, ?)", user.Email, user.Username, passwordHash, user.FirstName, user.LastName)
	return err
}

func (s *UserService) GetUsers() ([]GetUserResponse, error) {
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

func (s *UserService) GetUserByIdentifier(identifier string) (UserDB, error) {
	var user UserDB
	err := s.db.QueryRow("SELECT id, username, email, password_hash, first_name, last_name FROM users WHERE email = ? OR username = ?", identifier, identifier).Scan(&user.ID, &user.Username, &user.Email, &user.PasswordHash, &user.FirstName, &user.LastName)
	return user, err
}

func (s *UserService) getUserById(id int) (UserDB, error) {
	var user UserDB
	err := s.db.QueryRow("SELECT id, username, email, password_hash, first_name, last_name FROM users WHERE id = ?", id).Scan(&user.ID, &user.Username, &user.Email, &user.PasswordHash, &user.FirstName, &user.LastName)
	return user, err
}
