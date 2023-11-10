package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func New(addr string, BaseURL string) *Api {
	return &Api{
		addr:    addr,
		BaseURL: BaseURL,
		urlMap:  make(map[string]string),
	}
}

func (s *Api) Start() error {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.POST("/", s.ShortenURLHandler)
	r.GET("/:id", s.RedirectToOriginalURLHandler)

	err := r.Run(s.addr)
	if err != nil {
		fmt.Println("failed to start the browser")
		return err
	}

	return nil
}
