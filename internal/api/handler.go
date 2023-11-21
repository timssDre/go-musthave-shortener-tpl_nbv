package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"

	"net/http"
	"strings"
)

type StructURL struct {
	Performance string `json:"url"`
}

func (s *RestAPI) ShortenURLHandler(c *gin.Context) {
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to read request body", http.StatusInternalServerError)
		return
	}
	URLtoBody := strings.TrimSpace(string(body))
	shortURL := s.StructService.GetShortURL(URLtoBody)
	StructPerformance := StructURL{Performance: shortURL}
	respJSON, err := json.Marshal(StructPerformance)
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to read request body", http.StatusInternalServerError)
		return
	}
	c.Header("Content-Type", "application/json")
	c.Data(http.StatusCreated, "application/json", respJSON)
}

func (s *RestAPI) RedirectToOriginalURLHandler(c *gin.Context) {
	shortID := c.Param("id")
	originalURL, exists := s.StructService.GetOriginalURL(shortID)
	if !exists {
		c.String(http.StatusTemporaryRedirect, "URL not found")
		return
	}
	c.Header("Location", originalURL)
	c.String(http.StatusTemporaryRedirect, originalURL)
}
