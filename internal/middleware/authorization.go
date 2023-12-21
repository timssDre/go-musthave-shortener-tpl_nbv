package middleware

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

func AuthorizationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, err := getUserIDFromCookie(c)
		if err != nil {
			code := http.StatusBadRequest
			if errors.Is(err, http.ErrNoCookie) {
				code = http.StatusUnauthorized
			}
			contentType := c.Request.Header.Get("Content-Type")
			if contentType == "text/plain" {
				c.String(code, "Unauthorized")
			}
			if contentType == "application/json" {
				errorMassage := map[string]interface{}{
					"message": "Unauthorized",
					"code":    code,
				}
				answer, _ := json.Marshal(errorMassage)
				c.Data(code, "application/json", answer)
			}
			c.Abort()
			return
		}
		c.Set("userID", userID)
	}
}

func getUserIDFromCookie(ctx *gin.Context) (string, error) {
	userID, err := ctx.Cookie("userID")
	if err != nil {
		newUserID := uuid.New().String()
		ctx.SetCookie("userID", newUserID, 3600, "/", "localhost", false, true)
		return "", err
	}

	return userID, nil
}
