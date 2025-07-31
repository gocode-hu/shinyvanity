package main

import (
	"fmt"
	"github.com/gocode-hu/shinyvanity/internal/handler"
	"github.com/joho/godotenv"
	"net/http"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Cannot load .env file, using environment variables")
	}

	http.HandleFunc("/", handler.VanityHandler)
	listenPort := os.Getenv("LISTEN_PORT")
	if listenPort == "" {
		panic("LISTEN_PORT environment variable is not set")
	}
	address := ":" + listenPort

	fmt.Println("Listening on " + address)
	httpErr := http.ListenAndServe(address, nil)
	if httpErr != nil {
		panic(httpErr)
	}
}
