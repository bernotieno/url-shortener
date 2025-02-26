package main

import (
	"html/template"
	"log"
	"net/http"
	"url-shortener/handler"
)

func main() {
	fs := http.FileServer(http.Dir("frontend/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			handler.RedirectHandler(w, r)
			return
		}

		t := template.Must(template.ParseFiles("frontend/template/index.html"))
		t.Execute(w, nil)
	})
	http.HandleFunc("/api/shorten", handler.ShortenURLHandler)

	port := "8080"
	log.Printf("Server running at http://localhost:%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
