package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"

	"net/http"
	"strings"
)

type Request struct {
	URL string `json:"url"`
}

type Response struct {
	Result string `json:"result"`
}

func (s *RestAPI) ShortenURLHandler(c *gin.Context) {
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to read request body", http.StatusInternalServerError)
		return
	}
	url := strings.TrimSpace(string(body))
	shortURL := s.StructService.GetShortURL(url)
	if err = s.StructDump.Set(url, shortURL); err != nil {
		c.String(http.StatusInternalServerError, "failed to record event to file", http.StatusInternalServerError)
		return
	}

	c.Header("Content-Type", "text/plain")
	c.String(http.StatusCreated, shortURL)
}

func (s *RestAPI) ShortenURLJSON(c *gin.Context) {
	var decoderBody Request
	decoder := json.NewDecoder(c.Request.Body)
	err := decoder.Decode(&decoderBody)
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to read request body", http.StatusInternalServerError)
		return
	}
	url := strings.TrimSpace(decoderBody.URL)
	shortURL := s.StructService.GetShortURL(url)
	if err = s.StructDump.Set(url, shortURL); err != nil {
		c.String(http.StatusInternalServerError, "failed to record event to file", http.StatusInternalServerError)
		return
	}
	StructPerformance := Response{Result: shortURL}
	respJSON, err := json.Marshal(StructPerformance)
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to read request body", http.StatusInternalServerError)
		return
	}

	c.Header("Content-Type", "application/json")
	c.Data(http.StatusCreated, "application/json", respJSON)
}

func (s *RestAPI) RedirectToOriginalURL(c *gin.Context) {
	shortID := c.Param("id")
	originalURL, exists := s.StructService.GetOriginalURL(shortID)
	if !exists {
		c.String(http.StatusTemporaryRedirect, "URL not found")
		return
	}
	c.Header("Location", originalURL)
	c.String(http.StatusTemporaryRedirect, originalURL)
}
