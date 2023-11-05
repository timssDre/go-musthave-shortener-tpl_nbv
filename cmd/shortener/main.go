package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"math/rand"
	"net/http"
	"strings"
)

var urlMap = make(map[string]string)

func main() {
	r := gin.Default()
	r.POST("/", shortenURLHandler)
	r.GET("/:id", redirectToOriginalURLHandler)
	err := r.Run(":8080")
	if err != nil {
		fmt.Println("failed to start the browser")
	}
}

func shortenURLHandler(c *gin.Context) {
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to read request body", http.StatusInternalServerError)
		return
	}
	URLtoBody := strings.TrimSpace(string(body))

	shortID := randSeq(8)
	urlMap[shortID] = URLtoBody

	shortURL := fmt.Sprintf("http://localhost:8080/%s", shortID)

	c.Header("Content-Type", "text/plain")
	c.String(http.StatusCreated, shortURL)
}

func redirectToOriginalURLHandler(c *gin.Context) {
	shortID := c.Param("id")
	fmt.Println(shortID)
	originalURL, exists := urlMap[shortID]
	if exists {
		c.Header("Location", originalURL)
		c.String(http.StatusTemporaryRedirect, originalURL)
	} else {
		c.String(http.StatusTemporaryRedirect, "URL not found")
	}
}

func randSeq(n int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
