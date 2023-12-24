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

type RequestBodyURLs struct {
	CorrelationID string `json:"correlation_id"`
	OriginalURL   string `json:"original_url"`
}

type ResponseBodyURLs struct {
	CorrelationID string `json:"correlation_id"`
	ShortURL      string `json:"short_url"`
}

func (s *RestAPI) ShortenURLHandler(c *gin.Context) {
	httpStatus := http.StatusCreated
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to read request body", http.StatusInternalServerError)
		return
	}
	userIDFromContext, _ := c.Get("userID")
	userID, _ := userIDFromContext.(string)
	s.StructService.UserID = userID

	url := strings.TrimSpace(string(body))
	shortURL, err := s.StructService.Set(url)
	if err != nil {
		shortURL, err = s.StructService.GetExistURL(url, err)
		if err != nil {
			c.String(http.StatusInternalServerError, "the url could not be shortened", http.StatusInternalServerError)
			return
		}
		httpStatus = http.StatusConflict
	}
	c.Header("Content-Type", "text/plain")
	c.String(httpStatus, shortURL)
}

func (s *RestAPI) ShortenURLJSON(c *gin.Context) {
	var decoderBody Request
	httpStatus := http.StatusCreated
	decoder := json.NewDecoder(c.Request.Body)
	err := decoder.Decode(&decoderBody)
	c.Header("Content-Type", "application/json")
	if err != nil {
		errorMassage := map[string]interface{}{
			"message": "Failed to read request body",
			"code":    http.StatusInternalServerError,
		}
		answer, _ := json.Marshal(errorMassage)
		c.Data(http.StatusInternalServerError, "application/json", answer)
		return
	}

	userIDFromContext, _ := c.Get("userID")
	userID, _ := userIDFromContext.(string)
	s.StructService.UserID = userID

	url := strings.TrimSpace(decoderBody.URL)
	shortURL, err := s.StructService.Set(url)
	if err != nil {
		shortURL, err = s.StructService.GetExistURL(url, err)
		if err != nil {
			errorMassage := map[string]interface{}{
				"message": "the url could not be shortened",
				"code":    http.StatusInternalServerError,
			}
			answer, _ := json.Marshal(errorMassage)
			c.Data(http.StatusInternalServerError, "application/json", answer)
			return
		}
		httpStatus = http.StatusConflict
	}

	StructPerformance := Response{Result: shortURL}
	respJSON, err := json.Marshal(StructPerformance)
	if err != nil {
		errorMassage := map[string]interface{}{
			"message": "Failed to read request body",
			"code":    http.StatusInternalServerError,
		}
		var answer []byte
		answer, err = json.Marshal(errorMassage)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
			return
		}
		c.Data(http.StatusInternalServerError, "application/json", answer)
		return
	}
	c.Data(httpStatus, "application/json", respJSON)
}

func (s *RestAPI) RedirectToOriginalURL(c *gin.Context) {
	shortID := c.Param("id")
	originalURL, exists := s.StructService.Get(shortID)
	if !exists {
		c.String(http.StatusTemporaryRedirect, "URL not found")
		return
	}
	c.Header("Location", originalURL)
	c.String(http.StatusTemporaryRedirect, originalURL)
}

func (s *RestAPI) ShortenURLsJSON(c *gin.Context) {
	var decoderBody []RequestBodyURLs
	httpStatus := http.StatusCreated
	decoder := json.NewDecoder(c.Request.Body)
	err := decoder.Decode(&decoderBody)
	c.Header("Content-Type", "application/json")
	if err != nil {
		errorMassage := map[string]interface{}{
			"message": "Failed to read request body",
			"code":    http.StatusInternalServerError,
		}
		var answer []byte
		answer, err = json.Marshal(errorMassage)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
			return
		}
		c.Data(http.StatusInternalServerError, "application/json", answer)
		return
	}

	userIDFromContext, _ := c.Get("userID")
	userID, _ := userIDFromContext.(string)
	s.StructService.UserID = userID

	var URLResponses []ResponseBodyURLs
	for _, req := range decoderBody {
		url := strings.TrimSpace(req.OriginalURL)
		shortURL, err := s.StructService.Set(url)
		if err != nil {
			shortURL, err = s.StructService.GetExistURL(url, err)
			if err != nil {
				errorMassage := map[string]interface{}{
					"message": "the url could not be shortened",
					"code":    http.StatusInternalServerError,
				}
				var answer []byte
				answer, err = json.Marshal(errorMassage)
				if err != nil {
					c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
					return
				}
				c.Data(http.StatusInternalServerError, "application/json", answer)
				return
			}
			httpStatus = http.StatusConflict
		}
		urlResponse := ResponseBodyURLs{
			req.CorrelationID,
			shortURL,
		}
		URLResponses = append(URLResponses, urlResponse)
	}
	respJSON, err := json.Marshal(URLResponses)
	if err != nil {
		errorMassage := map[string]interface{}{
			"message": "Failed to read request body",
			"code":    http.StatusInternalServerError,
		}
		var answer []byte
		answer, err = json.Marshal(errorMassage)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
			return
		}
		c.Data(http.StatusInternalServerError, "application/json", answer)
		return
	}
	c.Data(httpStatus, "application/json", respJSON)

}

func (s *RestAPI) Ping(ctx *gin.Context) {
	err := s.StructService.Ping()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "")
		return
	}
	ctx.JSON(http.StatusOK, "")
}

func (s *RestAPI) UserURLsHandler(ctx *gin.Context) {
	code := http.StatusOK
	userIDFromContext, _ := ctx.Get("userID")
	userID, _ := userIDFromContext.(string)
	s.StructService.UserID = userID
	urls, err := s.StructService.GetFullRep()
	if err != nil {
		code = http.StatusInternalServerError
		errorMassage := map[string]interface{}{
			"message": "Failed to retrieve user URLs",
			"code":    code,
		}
		var answer []byte
		answer, err = json.Marshal(errorMassage)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
			return
		}
		ctx.Data(code, "application/json", answer)
		return
	}
	if len(urls) == 0 {
		code = http.StatusNoContent
		var errorMassages []map[string]interface{}
		errorMassage := map[string]interface{}{
			"message": "No URLs found",
			"code":    code,
		}
		errorMassages = append(errorMassages, errorMassage)
		var answer []byte
		answer, err = json.Marshal(errorMassages)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
			return
		}
		ctx.Data(code, "application/json", answer)

		return
	}

	respJSON, err := json.Marshal(urls)
	if err != nil {
		code = http.StatusInternalServerError
		var errorMassages []map[string]interface{}
		errorMassage := map[string]interface{}{
			"message": "Failed to marshal response",
			"code":    code,
		}
		errorMassages = append(errorMassages, errorMassage)
		var answer []byte
		answer, err = json.Marshal(errorMassages)
		if err != nil {
			ctx.AbortWithStatusJSON(code, gin.H{"message": "Internal Server Error"})
			return
		}
		ctx.Data(code, "application/json", answer)
		return
	}

	ctx.Data(code, "application/json", respJSON)
}
