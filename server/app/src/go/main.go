package main

import (
	"log"
	"net/http"

	"software-slayer/auth"
	"software-slayer/db"
	_ "software-slayer/docs"
	"software-slayer/learnings"
	"software-slayer/user"

	httpSwagger "github.com/swaggo/http-swagger"
)

/*
 * Manage the database connection and start the server
 */
func main() {
	auth.Init()

	db.Open()
	defer db.Close()

	http.Handle("/swagger/", httpSwagger.WrapHandler)

	user.InitUserRoutes()
	learnings.InitLearningRoutes()

	log.Println("Starting server on port 8080")
	http.ListenAndServe(":8080", nil)
}
