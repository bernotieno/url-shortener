package handler

import (
	"html/template"
	"net/http"
)

type ErrorData struct {
	StatusCode int
	Message    string
}

func RenderErrorPage(w http.ResponseWriter, statusCode int, message string) {
	tmpl, err := template.ParseFiles("frontend/template/errorpage.html")
	if err != nil {
		http.Error(w, "Error loading the error page", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(statusCode)
	tmpl.Execute(w, ErrorData{StatusCode: statusCode, Message: message})
}
