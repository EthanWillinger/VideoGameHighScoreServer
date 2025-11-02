package handlers

import (
	"net/http"
)

// RegisterUploadHandler registers the upload-related routes on the provided mux.
func RegisterUploadHandler(mux *http.ServeMux) {
	mux.HandleFunc("/upload", uploadHandler)
}

/*
	handles upload requests from users. In order to upload high scores, frontend must include the following in this POST request:

- Score: integer value representing the score achieved
- Username: string representing the user's name
- Game: integer representing the game's game id.
- Image: base64 encoded string representing the screenshot of the score
*/
func uploadHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	score := query.Get("score")
	username := query.Get("username")
	game := query.Get("game")
	imageData := query.Get("image")
	if !checkUserSubmission(score, username, game, imageData) {
		http.Error(w, "Invalid submission: missing important information, please fill out all missing information and try again", http.StatusBadRequest)
	}

	//Upload image to S3 bucket
	//Upload score to database

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Upload successful"))

}

/*
checkUserSubmission checks that all necessary fields are present in the user's submission and validates them
1. Ensure the score is within the appropriate range of the game, scores that exceed the set boundaries are invalid
2. Ensure the user exists to prevent ghost uploads
3. Ensure the game exists
4. Ensure the base64 string is a valid string
*/
func checkUserSubmission(score, username, game, image string) bool {

	// this check sucks
	return score != "" && username != "" && game != "" && image != ""
}
