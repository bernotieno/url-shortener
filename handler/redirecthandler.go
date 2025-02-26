package handler

import (
	"net/http"
	"strings"
)

func RedirectHandler(w http.ResponseWriter, r *http.Request) {
	shortKey := strings.TrimPrefix(r.URL.Path, "/")

	mutex.RLock()
	longURL, exists := urlStore[shortKey]
	mutex.RUnlock()

	if !exists {
		RenderErrorPage(w, http.StatusNotFound, "Short URL not found")
		return
	}

	http.Redirect(w, r, longURL, http.StatusFound)
}
