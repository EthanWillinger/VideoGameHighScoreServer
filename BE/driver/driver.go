/*
This code handles db actions and the initial connection
*/
package driver

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	objects "videogamehighscoreserver/objects"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func InitializeDB() (dberr error) {
	db, err := sql.Open("pgx", os.Getenv("PG_DSN"))
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
		return err
	}
	defer db.Close()
	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatalf("Unable to reach database: %v\n", pingErr)
		return pingErr
	}
	fmt.Println("Connected to database!")
	return nil
}

// Query to get user data from username
func GetUser(username string, db *sql.DB) (user objects.User, exists bool, err error) {
	err = db.QueryRow("SELECT id, username, email, password FROM users WHERE username=$1", 1).Scan(&user.ID, &user.Username, &user.Email, &user.Passkey)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, false, nil
		}
		return user, false, err
	}
	return user, true, nil
}

// Query to game data from game id
func GetGame(gameID int, db *sql.DB) (game objects.Game, exists bool, err error) {
	err = db.QueryRow("SELECT score_id, title, platform, version, release_date FROM games WHERE score_id=$1", gameID).Scan(&game.ID, &game.Title, &game.Platform, &game.Version, &game.ReleaseDate)
	if err != nil {
		if err == sql.ErrNoRows {
			return game, false, nil
		}
		return game, false, err
	}
	return game, true, nil
}

// Query to get scores from game id
func GetScoresByGame(gameID int, db *sql.DB) (scores []objects.Score, exists bool, err error) {
	rows, err := db.Query("SELECT id, game_id, user_id, score, date, score_type FROM scores WHERE game_id=$1 ORDER BY score DESC", gameID)
	if err != nil {
		return nil, false, err
	}
	defer rows.Close()
	for rows.Next() {
		var score objects.Score
		err := rows.Scan(&score.ID, &score.GameID, &score.UserID, &score.Score, &score.Date, &score.ScoreType)
		if err != nil {
			return nil, false, err
		}
		scores = append(scores, score)
	}
	return scores, true, nil
}
