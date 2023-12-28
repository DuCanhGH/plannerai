package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"planner/api"
	"planner/middleware"

	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	server := http.NewServeMux()

	server.Handle("/getques", middleware.Middleware(api.GetQuesHandler()))
	server.Handle("/getplan", middleware.Middleware(api.GetPlanHandler()))

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{os.Getenv("CLIENT_URL")},
		AllowCredentials: true,
		AllowedMethods:   []string{"POST"},
	})

	handler := c.Handler(server)

	fmt.Printf("Listening at port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, handler))
}
