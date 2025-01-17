package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"software-slayer/auth"
	"software-slayer/db"
	_ "software-slayer/docs"
	"software-slayer/learnings"
	"software-slayer/user"

	httpSwagger "github.com/swaggo/http-swagger"
)

const (
	TOKEN_LIFETIME          = time.Hour * 24
	JWT_SECRET_FILE_ENV_VAR = "JWT_SECRET_FILE"
)

/*
 * Manage the database connection and start the server
 */
func main() {
	tokenService := initTokenService()

	db, err := initDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	initSwagger()

	user.InitUserRest(user.NewUserService(db), tokenService)
	learnings.InitLearningsRest(learnings.NewLearningsService(db), tokenService)

	startServer()
}

/*
 * Initialize the auth package with the jwt secret and token lifetime
 */
func initTokenService() *auth.TokenService {
	jwtSecret, err := os.ReadFile(os.Getenv(JWT_SECRET_FILE_ENV_VAR))
	if err != nil {
		log.Fatal("Failed to read jwt secret file")
	}
	return auth.NewTokenService(TOKEN_LIFETIME, jwtSecret)
}

/*
 * Initialize the database connection
 */
func initDB() (*db.Database, error) {
	password, err := os.ReadFile(os.Getenv("DB_PASSWORD_FILE"))
	if err != nil {
		log.Fatal(err)
	}

	user := os.Getenv("DB_USER")
	address := os.Getenv("DB_ADDRESS")
	name := os.Getenv("DB_NAME")

	return db.NewDB(user, string(password), address, name)
}

/*
 * Initialize the swagger documentation
 */
func initSwagger() {
	http.Handle("/swagger/", httpSwagger.WrapHandler)
}

/*
 * Start the server
 */
func startServer() {
	log.Println("Starting server on port 8080")
	http.ListenAndServe(":8080", nil)
}
