package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"planner/api"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	api.GetQuesHandler()

	fmt.Printf("Listening at port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
