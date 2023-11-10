package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"math/rand"
	"net/http"
	"strings"
)

func (s *Api) ShortenURLHandler(c *gin.Context) {
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to read request body", http.StatusInternalServerError)
		return
	}
	URLtoBody := strings.TrimSpace(string(body))

	shortID := randSeq(8)
	s.urlMap[shortID] = URLtoBody

	shortURL := fmt.Sprintf("%s/%s", s.BaseURL, shortID)

	c.Header("Content-Type", "text/plain")
	c.String(http.StatusCreated, shortURL)
}

func randSeq(n int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func (s *Api) RedirectToOriginalURLHandler(c *gin.Context) {
	shortID := c.Param("id")
	originalURL, exists := s.urlMap[shortID]
	if exists {
		c.Header("Location", originalURL)
		c.String(http.StatusTemporaryRedirect, originalURL)
	} else {
		c.String(http.StatusTemporaryRedirect, "URL not found")
	}
}
