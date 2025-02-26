package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
	"url-shortener/handler"
)

func port() int16 {
	var port int16 = 9000

	portStr, found := os.LookupEnv("PORT")
	if !found {
		return port
	}
	iport, err := strconv.Atoi(portStr)
	if err != nil {
		return port
	}
	return int16(iport)
}

func main() {
	port := port()
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

	p := fmt.Sprintf(":%d", port)
	fmt.Println("Server started at ", p)
	log.Fatal(http.ListenAndServe(p, nil))
}
