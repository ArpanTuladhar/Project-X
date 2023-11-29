// routes/routes.go

package routes

import (
	"github.com/gorilla/mux"
	"github.com/yourusername/twitter-clone/controllers"
)

func SetupRoutes(r *mux.Router) {
	// Existing routes...

	// Authentication routes
	r.HandleFunc("/register", controllers.Register).Methods("POST")
	r.HandleFunc("/login", controllers.Login).Methods("POST")

	// Tweet routes
	r.HandleFunc("/create_tweet", controllers.CreateTweet).Methods("POST")
	r.HandleFunc("/tweets", controllers.GetTweets).Methods("GET")
	r.HandleFunc("/edit_tweet/{id}", controllers.EditTweet).Methods("PUT")
	r.HandleFunc("/delete_tweet/{id}", controllers.DeleteTweet).Methods("DELETE")
	r.HandleFunc("/like_tweet/{id}", controllers.LikeTweet).Methods("POST")
	r.HandleFunc("/comment_tweet/{id}", controllers.CommentTweet).Methods("POST")
	r.HandleFunc("/comments/{id}", controllers.GetComments).Methods("GET")
}
