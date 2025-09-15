package main

import (
	"log"
	"os"

	"github.com/srbhgalinde/url-shortner/internal/http"
)

func main() {
	router := http.NewRouter()

	portno := os.Getenv("PORT")
	if portno == "" {
		portno = ":8080"
	}

	log.Println("Server running on port ", portno)
	if err := router.Run(portno); err != nil {
		log.Fatal(err)
	}
}
