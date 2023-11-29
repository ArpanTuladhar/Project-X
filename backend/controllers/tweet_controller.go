// backend/controllers/tweet_controller.go

package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Tweet represents a tweet structure
type Tweet struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
	Likes   int    `json:"likes"`
}

// Comment represents a comment structure
type Comment struct {
	ID            int    `json:"id"`
	TweetID       int    `json:"tweet_id"`
	CommenterName string `json:"commenter_name"`
	Content       string `json:"content"`
}

// CreateTweet handles the creation of a new tweet
func CreateTweet(w http.ResponseWriter, r *http.Request) {
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

// GetTweets retrieves all tweets
func GetTweets(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id, content, likes FROM tweets")
	if err != nil {
		log.Println("Error fetching tweets from the database:", err)
		http.Error(w, "Error fetching tweets", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var tweets []Tweet

	for rows.Next() {
		var tweet Tweet
		err := rows.Scan(&tweet.ID, &tweet.Content, &tweet.Likes)
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

// EditTweet handles the editing of a tweet
func EditTweet(w http.ResponseWriter, r *http.Request) {
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

// DeleteTweet handles the deletion of a tweet
func DeleteTweet(w http.ResponseWriter, r *http.Request) {
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

// LikeTweet handles the liking of a tweet
func LikeTweet(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		vars := mux.Vars(r)
		tweetID := vars["id"]

		_, err := db.Exec("UPDATE tweets SET likes = likes + 1 WHERE id = ?", tweetID)
		if err != nil {
			log.Println("Error updating likes in the database:", err)
			http.Error(w, "Error updating likes", http.StatusInternalServerError)
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

// CommentTweet handles the commenting on a tweet
func CommentTweet(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		vars := mux.Vars(r)
		tweetID := vars["id"]

		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Unable to parse form data", http.StatusInternalServerError)
			return
		}

		commentContent := r.FormValue("commentContent")
		commenterName := "John Doe" // For simplicity, assuming commenter_name is fixed for now

		_, err = db.Exec("INSERT INTO comments (tweet_id, commenter_name, content) VALUES (?, ?, ?)", tweetID, commenterName, commentContent)
		if err != nil {
			log.Println("Error adding comment to the database:", err)
			http.Error(w, fmt.Sprintf("Error adding comment to tweet with ID %s: %v", tweetID, err), http.StatusInternalServerError)
			return
		}

		response := map[string]string{
			"status":  "success",
			"message": fmt.Sprintf("Comment added to tweet with ID: %s", tweetID),
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// GetComments retrieves comments for a specific tweet
func GetComments(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tweetID := vars["id"]

	rows, err := db.Query("SELECT id, tweet_id, commenter_name, content FROM comments WHERE tweet_id = ?", tweetID)
	if err != nil {
		log.Println("Error fetching comments from the database:", err)
		http.Error(w, "Error fetching comments", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var comments []Comment

	for rows.Next() {
		var comment Comment
		err := rows.Scan(&comment.ID, &comment.TweetID, &comment.CommenterName, &comment.Content)
		if err != nil {
			log.Println("Error scanning comment rows:", err)
			http.Error(w, "Error fetching comments", http.StatusInternalServerError)
			return
		}
		comments = append(comments, comment)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(comments)
}
