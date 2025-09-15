package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Sajid416/todo-app/cmd"
	"github.com/Sajid416/todo-app/database"
	"github.com/Sajid416/todo-app/router"
)

func main() {
	cmd.Serve()
	err := database.Connect()
	if err != nil {
		log.Println("Error in Database Connection")
	}
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}

	mux := router.RegisterRoutes()

	fmt.Println("Server running at:" + port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
	defer database.DB.Close()

}
