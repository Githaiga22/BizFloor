package handlers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func HandleErrors(w http.ResponseWriter, code int) {
	var message string
	switch code {
	case http.StatusNotFound:
		message = "Not Found"
	case http.StatusBadRequest:
		message = "Bad Request"
	case http.StatusMethodNotAllowed:
		message = "Method Not Allowed"
	case http.StatusForbidden:
		message = "Forbidden"
	default:
		message = "Internal Server Error"
	}

	data := struct {
		Title   string
		Status  int
		Message string
	}{
		Title:   "Error",
		Status:  code,
		Message: message,
	}

	// Set HTTP response status code
	w.WriteHeader(code)

	// Attempt to parse and execute the error template
	tmpl, err := template.ParseFiles("templates/error.html")
	if err != nil {
		// Log the error and provide fallback response
		log.Printf("Error parsing template: %v", err)
		serveFallbackErrorResponse(w, code, message)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		// Log the error and provide fallback response
		log.Printf("Error executing template: %v", err)
		serveFallbackErrorResponse(w, code, message)
	}
}

// serveFallbackErrorResponse provides a plain text fallback response in case of failure.
func serveFallbackErrorResponse(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	_, err := fmt.Fprintf(w, "%d - %s", code, message)
	if err != nil {
		log.Printf("Error writing fallback response: %v", err)
	}
}
