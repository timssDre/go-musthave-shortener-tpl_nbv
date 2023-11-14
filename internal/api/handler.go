package api

import (
	"github.com/gin-gonic/gin"
	"io"

	"net/http"
	"strings"
)

func (s *StructAPI) ShortenURLHandler(c *gin.Context) {
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to read request body", http.StatusInternalServerError)
		return
	}
	URLtoBody := strings.TrimSpace(string(body))

	shortURL := s.StructService.GetShortURL(URLtoBody)

	c.Header("Content-Type", "text/plain")
	c.String(http.StatusCreated, shortURL)
}

func (s *StructAPI) RedirectToOriginalURLHandler(c *gin.Context) {
	shortID := c.Param("id")
	originalURL, exists := s.StructService.GetOriginalURL(shortID)
	if !exists {
		c.String(http.StatusTemporaryRedirect, "URL not found")
		return
	}
	c.Header("Location", originalURL)
	c.String(http.StatusTemporaryRedirect, originalURL)
}
