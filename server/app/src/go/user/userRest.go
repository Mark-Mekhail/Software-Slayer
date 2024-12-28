package user

import (
	"encoding/json"
	"fmt"
	"net/http"

	"software-slayer/auth"
	"software-slayer/utils"
)

// @Summary Create a new user
// @Description Register a new user with an email, password, and name
// @Tags Users
// @Accept json
// @Param user body CreateUserRequest true "User object that needs to be added"
// @Success 201
// @Router /user [post]
func createUser(w http.ResponseWriter, r *http.Request) {
	var user CreateUserRequest
	if err := utils.Decode(w, r, &user); err != nil {
		return
	}

	passwordHash, err := auth.HashPassword(user.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = createUserDB(&user, passwordHash)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// @Summary Login
// @Description Login with an email and password
// @Tags Users
// @Accept json
// @Produce json
// @Param credentials body Credentials true "Credentials object that needs to be added"
// @Success 200 {object} map[string]string
// @Router /login [post]
func handleLogin(w http.ResponseWriter, r *http.Request) {
	var credentials Credentials
	if err := utils.Decode(w, r, &credentials); err != nil {
		return
	}

	user, err := getUserByIdentifierDB(credentials.Identifier)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = auth.ValidatePassword(credentials.Password, user.PasswordHash)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	token, err := auth.GenerateToken(user.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

// @Summary Get users
// @Description Get a filtered set of users
// @Tags Users
// @Produce json
// @Param Authorization header string false "Bearer token"
// @Param current query bool false "Get only the current user"
// @Success 200 {array} GetUserResponse
// @Router /user [get]
func getUsers(w http.ResponseWriter, r *http.Request) {
	current := r.URL.Query().Get("current")

	if current == "true" {
		getCurrentUser(w, r)
	} else {
		getAllUsers(w, r)
	}
}

func getCurrentUser(w http.ResponseWriter, r *http.Request) {
	userId, err := auth.AuthorizeUser(r)

	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	user, err := getUserByIdDB(userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	currentUserResponse := GetCurrentUserResponse{
		Email: user.Email,
		GetUserResponse: GetUserResponse{
			ID:       user.ID,
			UserBase: user.UserBase,
		},
	}
	json.NewEncoder(w).Encode(currentUserResponse)
}

func getAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := getUsersDB()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(users)
}

func handleUser(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		createUser(w, r)
	case http.MethodGet:
		getUsers(w, r)
	default:
		http.Error(w, fmt.Sprintf("Method %s not allowed", r.Method), http.StatusMethodNotAllowed)
	}
}

func InitUserRoutes() {
	http.HandleFunc("/user", handleUser)
	http.HandleFunc("/login", handleLogin)
}
