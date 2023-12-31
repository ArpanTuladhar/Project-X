package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("mysql", "root:Xonen@3616@tcp(127.0.0.1:3306)/twitter_clone")
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}

	log.Println("Connected to the database")
}

type Tweet struct {
	ID            int    `json:"id"`
	Content       string `json:"content"`
	LikesCount    int    `json:"likesCount"`
	CommentsCount int    `json:"commentsCount"`
}

func handleCreateTweet(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Unable to parse form data", http.StatusInternalServerError)
			return
		}

		tweetContent := r.FormValue("tweetContent")

		_, err = db.Exec("INSERT INTO tweets (content) VALUES (?)", tweetContent)
		if err != nil {
			log.Println("Error storing tweet in the database:", err)
			http.Error(w, "Error storing tweet in the database", http.StatusInternalServerError)
			return
		}

		response := map[string]string{
			"status":  "success",
			"message": fmt.Sprintf("Tweet created: %s", tweetContent),
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func handleGetTweets(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id, content, likes_count, comments_count FROM tweets")
	if err != nil {
		log.Println("Error fetching tweets from the database:", err)
		http.Error(w, "Error fetching tweets", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var tweets []Tweet

	for rows.Next() {
		var tweet Tweet
		err := rows.Scan(&tweet.ID, &tweet.Content, &tweet.LikesCount, &tweet.CommentsCount)
		if err != nil {
			log.Println("Error scanning tweet rows:", err)
			http.Error(w, "Error fetching tweets", http.StatusInternalServerError)
			return
		}
		tweets = append(tweets, tweet)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tweets)
}

func handleEditTweet(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPut {
		vars := mux.Vars(r)
		tweetID := vars["id"]

		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Unable to parse form data", http.StatusInternalServerError)
			return
		}

		tweetContent := r.FormValue("tweetContent")

		_, err = db.Exec("UPDATE tweets SET content = ? WHERE id = ?", tweetContent, tweetID)
		if err != nil {
			log.Println("Error updating tweet in the database:", err)
			http.Error(w, "Error updating tweet", http.StatusInternalServerError)
			return
		}

		response := map[string]string{
			"status":  "success",
			"message": fmt.Sprintf("Tweet updated with ID: %s", tweetID),
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func handleDeleteTweet(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodDelete {
		vars := mux.Vars(r)
		tweetID := vars["id"]

		_, err := db.Exec("DELETE FROM tweets WHERE id = ?", tweetID)
		if err != nil {
			log.Println("Error deleting tweet:", err)
			http.Error(w, "Error deleting tweet", http.StatusInternalServerError)
			return
		}

		response := map[string]string{
			"status":  "success",
			"message": fmt.Sprintf("Tweet deleted with ID: %s", tweetID),
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func handleLikeTweet(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		vars := mux.Vars(r)
		tweetID := vars["id"]

		_, err := db.Exec("UPDATE tweets SET likes_count = likes_count + 1 WHERE id = ?", tweetID)
		if err != nil {
			log.Println("Error updating likes count for tweet:", err)
			http.Error(w, "Error updating likes count for tweet", http.StatusInternalServerError)
			return
		}

		response := map[string]string{
			"status":  "success",
			"message": fmt.Sprintf("Liked tweet with ID: %s", tweetID),
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func handleCommentTweet(w http.ResponseWriter, r *http.Request) {
	// Similar logic as handleLikeTweet, increment comments_count instead
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/create_tweet", handleCreateTweet).Methods("POST")
	r.HandleFunc("/tweets", handleGetTweets).Methods("GET")
	r.HandleFunc("/edit_tweet/{id}", handleEditTweet).Methods("PUT")
	r.HandleFunc("/delete_tweet/{id}", handleDeleteTweet).Methods("DELETE")
	r.HandleFunc("/like_tweet/{id}", handleLikeTweet).Methods("POST")
	r.HandleFunc("/comment_tweet/{id}", handleCommentTweet).Methods("POST")

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type"},
	})

	handler := c.Handler(r)

	http.Handle("/", handler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
