package user

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"software-slayer/auth"
	"software-slayer/utils"
)

var userService UserService
var tokenService auth.TokenService

// @Summary Create a new user
// @Description Register a new user with an email, password, and name
// @Tags Users
// @Accept json
// @Produce json
// @Param user body CreateUserRequest true "User object that needs to be added"
// @Success 201 {string} string "User created successfully"
// @Failure 400 {object} utils.ErrorResponse "Invalid request data"
// @Failure 500 {object} utils.ErrorResponse "Server error"
// @Router /user [post]
func createUser(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()
	
	var user CreateUserRequest
	if err := utils.Decode(w, r, &user); err != nil {
		return
	}

	log.Printf("Attempting to create user with email: %s, username: %s", user.Email, user.Username)

	if err := validateCreateUserRequest(user); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Invalid %s", err.Error()))
		return
	}

	passwordHash, err := auth.HashPassword(user.Password)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to process password")
		return
	}

	err = userService.CreateUser(ctx, &user, passwordHash)
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			utils.RespondWithError(w, http.StatusConflict, "A user with this email or username already exists")
			return
		}
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to create user")
		return
	}

	log.Printf("Successfully created user: %s", user.Username)
	utils.RespondWithJSON(w, http.StatusCreated, map[string]string{"message": "User created successfully"})
}

// @Summary Login
// @Description Login with an email/username and password
// @Tags Users
// @Accept json
// @Produce json
// @Param credentials body Credentials true "Credentials object for login"
// @Success 200 {object} LoginResponse
// @Failure 400 {object} utils.ErrorResponse "Invalid credentials format"
// @Failure 401 {object} utils.ErrorResponse "Authentication failed"
// @Failure 500 {object} utils.ErrorResponse "Server error"
// @Router /login [post]
func handleLogin(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()
	
	var credentials Credentials
	if err := utils.Decode(w, r, &credentials); err != nil {
		return
	}

	log.Printf("Login attempt for: %s", credentials.Identifier)

	user, err := userService.GetUserByIdentifier(ctx, credentials.Identifier)
	if err != nil {
		utils.RespondWithError(w, http.StatusUnauthorized, "Invalid credentials")
		return
	}

	err = auth.ValidatePassword(credentials.Password, user.PasswordHash)
	if err != nil {
		utils.RespondWithError(w, http.StatusUnauthorized, "Invalid credentials")
		return
	}

	token, err := tokenService.GenerateToken(user.ID)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to generate token")
		return
	}

	log.Printf("Successful login for user: %s (ID: %d)", user.Username, user.ID)

	loginResponse := LoginResponse{
		Token: token,
		UserInfo: GetCurrentUserResponse{
			Email: user.Email,
			GetUserResponse: GetUserResponse{
				ID:       user.ID,
				UserBase: user.UserBase,
			},
		},
	}

	utils.RespondWithJSON(w, http.StatusOK, loginResponse)
}

// @Summary Get users
// @Description Get a filtered set of users
// @Tags Users
// @Produce json
// @Param Authorization header string false "Bearer token"
// @Param current query bool false "Get only the current user"
// @Success 200 {array} GetUserResponse
// @Failure 400 {object} utils.ErrorResponse "Invalid request parameters"
// @Failure 401 {object} utils.ErrorResponse "Unauthorized"
// @Failure 500 {object} utils.ErrorResponse "Server error"
// @Router /user [get]
func getUsers(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()
	
	var current bool

	currentVal := r.URL.Query().Get("current")
	if currentVal == "" {
		current = false
	} else {
		var err error
		current, err = strconv.ParseBool(currentVal)
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid current parameter")
			return
		}
	}

	if current {
		getCurrentUser(ctx, w, r)
	} else {
		getAllUsers(ctx, w)
	}
}

/*
 * getCurrentUser gets the current user from the database and returns it as a response.
 * @param ctx: the request context
 * @param w: the response writer
 * @param r: the request
 */
func getCurrentUser(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	userId, err := tokenService.AuthorizeUser(r.Header.Get("Authorization"))

	if err != nil {
		utils.RespondWithError(w, http.StatusUnauthorized, "Invalid or missing authentication token")
		return
	}

	user, err := userService.GetUserById(ctx, userId)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to retrieve user information")
		return
	}

	log.Printf("Retrieved current user: %s (ID: %d)", user.Username, user.ID)

	currentUserResponse := GetCurrentUserResponse{
		Email: user.Email,
		GetUserResponse: GetUserResponse{
			ID:       user.ID,
			UserBase: user.UserBase,
		},
	}

	utils.RespondWithJSON(w, http.StatusOK, currentUserResponse)
}

/*
 * getAllUsers gets all users from the database and returns them as a response.
 * @param ctx: the request context
 * @param w: the response writer
 */
func getAllUsers(ctx context.Context, w http.ResponseWriter) {
	users, err := userService.GetUsers(ctx)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to retrieve users")
		return
	}

	log.Printf("Retrieved %d users", len(users))
	utils.RespondWithJSON(w, http.StatusOK, users)
}

// InitUserRest initializes the user REST endpoints
func InitUserRest(_userService UserService, _tokenService auth.TokenService) {
	userService = _userService
	tokenService = _tokenService

	http.HandleFunc("POST /user", createUser)
	http.HandleFunc("GET /user", getUsers)
	http.HandleFunc("POST /login", handleLogin)

	log.Println("User REST endpoints initialized")
}
