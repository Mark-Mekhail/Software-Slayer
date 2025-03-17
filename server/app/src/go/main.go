package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"software-slayer/auth"
	"software-slayer/configs"
	"software-slayer/db"
	_ "software-slayer/docs"
	"software-slayer/learnings"
	"software-slayer/user"

	httpSwagger "github.com/swaggo/http-swagger"
)

/*
 * Entry point for the application
 * Sets up the database connection, REST endpoints, and starts the server
 */
func main() {
	log.Println("Starting Software Slayer API server...")

	// Initialize services
	tokenService := initTokenService()
	database := initDB()
	defer database.Close()

	initSwagger()

	// Initialize REST handlers
	user.InitUserRest(user.NewUserService(database), tokenService)
	learnings.InitLearningsRest(learnings.NewLearningsService(database), tokenService)

	// Start server with graceful shutdown
	startServerWithGracefulShutdown()
}

/*
 * Initialize the auth package with the jwt secret and token lifetime
 */
func initTokenService() *auth.TokenServiceImpl {
	jwtSecretPath := os.Getenv(configs.JWT_SECRET_FILE_ENV_VAR)
	log.Printf("Reading JWT secret from: %s", jwtSecretPath)

	jwtSecret, err := os.ReadFile(jwtSecretPath)
	if err != nil {
		log.Fatalf("Failed to read JWT secret file: %v", err)
	}

	log.Println("JWT token service initialized successfully")
	return auth.NewTokenService(configs.TOKEN_LIFETIME, jwtSecret)
}

/*
 * Initialize the database connection
 */
func initDB() *db.Database {
	dbPasswordPath := os.Getenv("DB_PASSWORD_FILE")
	log.Printf("Reading database password from: %s", dbPasswordPath)

	password, err := os.ReadFile(dbPasswordPath)
	if err != nil {
		log.Fatalf("Failed to read database password: %v", err)
	}

	user := os.Getenv("DB_USER")
	address := os.Getenv("DB_ADDRESS")
	name := os.Getenv("DB_NAME")

	log.Printf("Connecting to database %s at %s as user %s", name, address, user)
	dbConn, err := db.OpenConnection(user, string(password), address, name)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	return db.NewDB(dbConn)
}

/*
 * Initialize the swagger documentation
 */
func initSwagger() {
	log.Println("Initializing Swagger documentation")
	http.Handle("/swagger/", httpSwagger.WrapHandler)
}

/*
 * Start the server with graceful shutdown capability
 * Handles OS signals to perform a clean shutdown
 */
func startServerWithGracefulShutdown() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Create a new server instance
	srv := &http.Server{
		Addr:         ":" + port,
		Handler:      http.DefaultServeMux,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Start server in a goroutine
	go func() {
		log.Printf("Server starting on port %s", port)
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("HTTP server error: %v", err)
		}
	}()

	// Set up channel for graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	// Block until signal is received
	<-stop
	log.Println("Shutting down server...")

	// Create context with timeout for graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Attempt to gracefully shutdown the server
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server shutdown failed: %v", err)
	}

	log.Println("Server gracefully stopped")
}
