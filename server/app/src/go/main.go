package main

import (
	"log"
	"net/http"

	"software-slayer/db"
	"software-slayer/user"
	"software-slayer/skills"
	_ "software-slayer/docs"

	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {
	db.Open()
	defer db.Close()

	http.Handle("/swagger/", httpSwagger.WrapHandler)
	
	user.InitUserRoutes()
	skills.InitSkillRoutes()

	log.Println("Starting server on port 8080")
	http.ListenAndServe(":8080", nil)
}