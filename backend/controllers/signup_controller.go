package controllers

import (
	"html/template"
	"net/http"
)

// SignupController handles requests for the signup page
func SignupController(w http.ResponseWriter, r *http.Request) {
	// Check if it's a POST request to process form submission
	if r.Method == http.MethodPost {
		// Retrieve form data
		username := r.FormValue("username")
		password := r.FormValue("password")

		// Implement signup logic (replace with your database logic)
		// Example: Insert the user into the database
		// _, err := db.Exec("INSERT INTO users (username, password) VALUES (?, ?)", username, password)
		// Check for errors and handle accordingly

		// For simplicity, we'll just print a success message for now
		w.Write([]byte("Signup successful!"))

		return
	}

	// If it's not a POST request, render the signup form
	renderTemplate(w, "signup.html", nil)
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
