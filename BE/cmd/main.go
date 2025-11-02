package main

import (
	"log"
	"net/http"
	"os"
	db "videogamehighscoreserver/driver"
	"videogamehighscoreserver/handlers"
)

func main() {
	mux := http.NewServeMux()

	//Initialize environment variables
	os.Setenv("PG_DSN", "user=postgres password=local host=localhost port=5432 dbname=vghighscore sslmode=disable")

	//Connect to database
	dberr := db.InitializeDB()
	if dberr != nil {
		log.Fatalf("Database initialization failed: %v", dberr)
	}

	// Register handlers from the handlers package.
	handlers.RegisterUploadHandler(mux)

	log.Println("starting server on :8080")
	log.Fatal(http.ListenAndServe(":8081", mux))
}
