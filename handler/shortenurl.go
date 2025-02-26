package handler

import (
	"encoding/json"
	"net/http"
	"strings"
)

func ShortenURLHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		RenderErrorPage(w, http.StatusMethodNotAllowed, "Invalid request method")
		return
	}

	var requestData struct {
		LongURL string `json:"url"`
	}
	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil || strings.TrimSpace(requestData.LongURL) == "" {
		RenderErrorPage(w, http.StatusBadRequest, "Invalid input")
		return
	}

	mutex.Lock()
	defer mutex.Unlock()

	if shortKey, exists := reverseStore[requestData.LongURL]; exists {
		json.NewEncoder(w).Encode(map[string]string{"short_url": "/" + shortKey})
		return
	}

	var shortKey string
	for {
		shortKey = generateShortKey()
		if _, exists := urlStore[shortKey]; !exists {
			break
		}
	}

	urlStore[shortKey] = requestData.LongURL
	reverseStore[requestData.LongURL] = shortKey

	json.NewEncoder(w).Encode(map[string]string{"short_url": "/" + shortKey})
}
