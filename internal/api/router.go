package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/timssDre/go-musthave-shortener-tpl_nbv.git/internal/storage"
)

func (s *storage.Storage) Start(ServerAddr string) error {
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
