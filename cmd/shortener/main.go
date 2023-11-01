package main

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

var urlMap = make(map[string]string)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handle)
	//mux.HandleFunc("/", shortenURLHandler)
	//mux.HandleFunc("/{id}", redirectToOriginalURLHandler)
	http.ListenAndServe(":8080", mux)
}

func handle(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		redirectToOriginalURLHandler(w, r)
	} else if r.Method == http.MethodPost {
		shortenURLHandler(w, r)
	}
}

func shortenURLHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method POST", http.StatusMethodNotAllowed)
		return
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}
	URLtoBody := strings.TrimSpace(string(body))

	shortID := randSeq(8)
	urlMap[shortID] = URLtoBody

	shortURL := fmt.Sprintf("http://localhost:8080/%s", shortID)

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, shortURL)
}

func redirectToOriginalURLHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method GET", http.StatusMethodNotAllowed)
		return
	}

	shortID := r.URL.Path[1:]

	fmt.Println(shortID)

	originalURL, exists := urlMap[shortID]
	if exists {
		w.Header().Set("Location", originalURL)
		http.Redirect(w, r, originalURL, http.StatusTemporaryRedirect)
	} else {
		http.Error(w, "URL not found", http.StatusBadRequest)
	}
}

func init() {
	rand.NewSource(time.Now().UnixNano())
}

func randSeq(n int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
