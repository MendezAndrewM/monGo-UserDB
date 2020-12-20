package main

import (
	"auth/routes"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	e := godotenv.Load()

	if e != nil {
		log.Fatal("Failed to load .env file")
	}
	fmt.Println(e)

	port := os.Getenv("PORT")

	http.Handle("/", routes.Handlers())

	log.Printf("Server running on Port#: '%s'", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
