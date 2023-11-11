package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func (s *Screwdriver) StartService(ServerAddr string) error {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.POST("/", s.ShortenURLHandler)
	r.GET("/:id", s.RedirectToOriginalURLHandler)

	err := r.Run(ServerAddr)
	if err != nil {
		fmt.Println("failed to start the browser")
		return err
	}

	return nil
}
