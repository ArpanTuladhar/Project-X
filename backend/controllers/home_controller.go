// controllers/home_controller.go

package controllers

import (
	"net/http"
)

// HomeController handles requests for the home page
func HomeController(w http.ResponseWriter, r *http.Request) {
	// Implement home page logic
	w.Write([]byte("Welcome to the home page!"))
}
