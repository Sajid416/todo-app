package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Sajid416/todo-app/database"
	"github.com/Sajid416/todo-app/router"
)

func main() {
	err := database.Connect()
	if err != nil {
		log.Println("Error in Database Connection")
	}
	err = os.Getenv("SERVER_PORT")

	mux := router.RegisterRoutes()
	fmt.Println()

}
