package user

type UserBase struct {
	Username  string `json:"username"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type UserDB struct {
	ID           int
	Email        string
	PasswordHash string
	UserBase
}

type GetUserResponse struct {
	ID 			 int `json:"id"`
	UserBase
}

type GetCurrentUserResponse struct {
	Email string `json:"email"`
	GetUserResponse
}

type CreateUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	UserBase
}

type Credentials struct {
	Identifier string `json:"identifier"`
	Password   string `json:"password"`
}
