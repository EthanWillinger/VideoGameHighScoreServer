package objects

import "time"

type Score struct {
	ID        int       `json:"id"`
	GameID    int       `json:"game_id"`
	UserID    int       `json:"user_id"`
	ImageName string    `json:"image_name"`
	Date      time.Time `json:"date"` // Use `time.Time` if you prefer
	Score     int       `json:"score"`
}
