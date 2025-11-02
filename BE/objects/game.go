package objects

import "time"

type Game struct {
	ID          int       `json:"score_id"`
	Title       string    `json:"title"`
	Platform    string    `json:"platform"`
	Version     string    `json:"version"`
	ReleaseDate time.Time `json:"release_date"`
}
