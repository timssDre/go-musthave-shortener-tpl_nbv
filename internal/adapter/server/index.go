package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
)

type Server struct {
	addr    string
	BaseURL string
	urlMap  map[string]string
}

func New(addr string, BaseURL string) *Server {
	return &Server{
		addr:    addr,
		BaseURL: BaseURL,
		urlMap:  make(map[string]string),
	}
}

func (s *Server) Start() error {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.POST("/", s.shortenURLHandler)
	r.GET("/:id", s.redirectToOriginalURLHandler)

	err := r.Run(s.addr)
	if err != nil {
		fmt.Println("failed to start the browser")
		return err
	}

	return nil
}

func randSeq(n int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
