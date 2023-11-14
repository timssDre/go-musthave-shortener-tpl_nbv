package api

import (
	"github.com/gin-gonic/gin"
)

func (s *StructAPI) setRoutes(r *gin.Engine) {
	r.POST("/", s.ShortenURLHandler)
	r.GET("/:id", s.RedirectToOriginalURLHandler)
}
