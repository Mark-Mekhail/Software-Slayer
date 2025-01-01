package user

import (
	"software-slayer/db"
)

/*
 * createUserDB creates a new user in the database
 * @param user: the user to be created
 * @param passwordHash: the hashed password of the user
 * @return error: an error if the user could not be created
 */
func createUserDB(user *CreateUserRequest, passwordHash string) error {
	_, err := db.Exec("INSERT INTO users (email, username, password_hash, first_name, last_name) VALUES (?, ?, ?, ?, ?)", user.Email, user.Username, passwordHash, user.FirstName, user.LastName)
	return err
}

/*
 * getUsersDB gets all users from the database
 * @return []GetUserResponse: a list of all users
 * @return error: an error if the users could not be retrieved
 */
func getUsersDB() ([]GetUserResponse, error) {
	rows, err := db.Query("SELECT id, username, first_name, last_name FROM users")
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

/*
 * getUserByIdentifierDB gets a user by their email or username
 * @param identifier: the email or username of the user
 * @return UserDB: the user with the given email or username
 * @return error: an error if the user could not be retrieved
 */
func getUserByIdentifierDB(identifier string) (UserDB, error) {
	var user UserDB
	err := db.QueryRow("SELECT id, username, email, password_hash, first_name, last_name FROM users WHERE email = ? OR username = ?", identifier, identifier).Scan(&user.ID, &user.Username, &user.Email, &user.PasswordHash, &user.FirstName, &user.LastName)
	return user, err
}

/*
 * getUserByIdDB gets a user by their id
 * @param id: the id of the user
 * @return UserDB: the user with the given id
 * @return error: an error if the user could not be retrieved
 */
func getUserByIdDB(id int) (UserDB, error) {
	var user UserDB
	err := db.QueryRow("SELECT id, username, email, password_hash, first_name, last_name FROM users WHERE id = ?", id).Scan(&user.ID, &user.Username, &user.Email, &user.PasswordHash, &user.FirstName, &user.LastName)
	return user, err
}
