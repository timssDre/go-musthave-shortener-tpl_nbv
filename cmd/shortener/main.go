package main

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", handle)
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		return
	}
}

func handle(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		redirectToOriginalURLHandler(w, r)
	} else if r.Method == http.MethodPost {
		shortenURLHandler(w, r)
	}
}

func shortenURLHandler(w http.ResponseWriter, r *http.Request) {
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
	shortID := r.URL.Path[1:]

	originalURL, exists := urlMap[shortID]
	if exists {
		w.Header().Set("Location", originalURL)
		http.Redirect(w, r, originalURL, http.StatusTemporaryRedirect)
	} else {
		http.Error(w, "URL not found", http.StatusBadRequest)
	}
}

}
