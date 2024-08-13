package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/PowerLightStar/WeatherForecast/routes"
	"github.com/joho/godotenv"
)

func main() {
	// Load the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Access environment variables
	port := os.Getenv("PORT")

	// assign the handler function
	routes.RegisterRoutes()

	fmt.Println("Server is Running....")
	// run http server on
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		fmt.Printf("Error starting server: %s\n", err)
		os.Exit(1)
	}
}
