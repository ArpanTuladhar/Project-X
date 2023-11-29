package models

type Tweet struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
	// Add other tweet-related fields as needed
}
