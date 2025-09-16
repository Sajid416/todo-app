package main

import (
	"log"

	"github.com/Sajid416/todo-app/cmd"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}
	cmd.Serve()

}
