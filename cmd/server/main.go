package main

import (
	"log"

	"github.com/srbhgalinde/url-shortner/internal/http"
)

func main() {
	router := http.NewRouter()

	// portno:=getenv()
	portno := ":8080"
	log.Println("Server running on port ", portno)
	if err := router.Run(portno); err != nil {
		log.Fatal(err)
	}
}
