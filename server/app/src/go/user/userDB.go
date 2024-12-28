package user

import (
	"software-slayer/db"
)

func createUserDB(user *CreateUserRequest, passwordHash string) error {
	_, err := db.Exec("INSERT INTO users (email, username, password_hash, first_name, last_name) VALUES (?, ?, ?, ?, ?)", user.Email, user.Username, passwordHash, user.FirstName, user.LastName)
	return err
}

func getUsersDB() ([]GetUserResponse, error) {
	rows, err := db.Query("SELECT id, username, first_name, last_name FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []GetUserResponse{}
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

// Identifier should be email or username
func getUserByIdentifierDB(identifier string) (UserDB, error) {
	var user UserDB
	err := db.QueryRow("SELECT id, username, email, password_hash, first_name, last_name FROM users WHERE email = ? OR username = ?", identifier, identifier).Scan(&user.ID, &user.Username, &user.Email, &user.PasswordHash, &user.FirstName, &user.LastName)
	return user, err
}

func getUserByIdDB(id int) (UserDB, error) {
	var user UserDB
	err := db.QueryRow("SELECT id, username, email, password_hash, first_name, last_name FROM users WHERE id = ?", id).Scan(&user.ID, &user.Username, &user.Email, &user.PasswordHash, &user.FirstName, &user.LastName)
	return user, err
}
