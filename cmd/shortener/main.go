package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/timssDre/go-musthave-shortener-tpl_nbv.git/cmd/shortener/config"
	"io"
	"math/rand"
	"net/http"
	"strings"
)

var (
	urlMap  = make(map[string]string)
	addr    string
	BaseUrl string
)

func main() {
	addrConfig := config.InitConfig()

	flag.StringVar(&addr, "a", addrConfig.ServerAddr, "address and port to run server")
	flag.StringVar(&BaseUrl, "b", addrConfig.BaseUrl, "address and port to run server addrResPos")
	flag.Parse()

	r := gin.Default()
	r.POST("/", shortenURLHandler)
	r.GET("/:id", redirectToOriginalURLHandler)

	err := r.Run(addr)
	if err != nil {
		fmt.Println("failed to start the browser")
		panic(err)
	}
}

func shortenURLHandler(c *gin.Context) {
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to read request body", http.StatusInternalServerError)
		return
	}
	URLtoBody := strings.TrimSpace(string(body))

	shortID := randSeq(8)
	urlMap[shortID] = URLtoBody

	shortURL := fmt.Sprintf("http://%s/%s", addr, shortID)

	c.Header("Content-Type", "text/plain")
	c.String(http.StatusCreated, shortURL)
}

func redirectToOriginalURLHandler(c *gin.Context) {
	shortID := c.Param("id")
	fmt.Println(shortID)
	originalURL, exists := urlMap[shortID]
	if exists {
		c.Header("Location", originalURL)
		c.String(http.StatusTemporaryRedirect, originalURL)
	} else {
		c.String(http.StatusTemporaryRedirect, "URL not found")
	}
}

func randSeq(n int) string {
	if BaseUrl != "default" {
		return BaseUrl
	}

	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
