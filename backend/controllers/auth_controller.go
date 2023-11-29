package controllers

import (
	"html/template"
	"net/http"
)

// LoginController handles requests for the login page
func LoginController(w http.ResponseWriter, r *http.Request) {
	// Check if it's a POST request to process form submission
	if r.Method == http.MethodPost {
		// Retrieve form data
		username := r.FormValue("username")
		password := r.FormValue("password")

		// Implement login logic (replace with your database logic)
		// Example: Retrieve the user from the database based on the username
		// storedUser, err := getUserByUsername(username)
		// Check for errors and handle accordingly

		// Compare hashed password from the database with the provided password
		// Example: err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(password))
		// Check for errors and handle accordingly

		// If credentials are valid, generate JWT token
		// token, err := GenerateToken(storedUser)
		// Check for errors and handle accordingly

		// For simplicity, we'll just print a success message for now
		w.Write([]byte("Login successful!"))

		return
	}

	// If it's not a POST request, render the login form
	renderTemplate(w, "login.html", nil)
}

// Helper function to render HTML templates
func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	tmpl, err := template.ParseFiles("templates/" + tmpl)
	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}
