package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"

	"net/http"
	"strings"
)

func (s *StructApi) ShortenURLHandler(c *gin.Context) {
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to read request body", http.StatusInternalServerError)
		return
	}
	URLtoBody := strings.TrimSpace(string(body))

	shortID := RandSeq(8)
	s.storage.SetValueMap(shortID, URLtoBody)
	shortURL := fmt.Sprintf("%s/%s", s.BaseURL, shortID)

	c.Header("Content-Type", "text/plain")
	c.String(http.StatusCreated, shortURL)
}

func (s *StructApi) RedirectToOriginalURLHandler(c *gin.Context) {
	shortID := c.Param("id")
	originalURL, exists := s.storage.GetValueMap(shortID)
	if exists {
		c.Header("Location", originalURL)
		c.String(http.StatusTemporaryRedirect, originalURL)
		return
	}
	c.String(http.StatusTemporaryRedirect, "URL not found")
}
