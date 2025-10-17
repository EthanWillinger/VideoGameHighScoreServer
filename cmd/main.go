package main

import (
	"log"
	"net/http"

	"videogamehighscoreserver/handlers"
)

func main() {
	mux := http.NewServeMux()

	// Register handlers from the handlers package.
	handlers.RegisterUploadHandler(mux)

	log.Println("starting server on :8080")
	log.Fatal(http.ListenAndServe(":8081", mux))
}
