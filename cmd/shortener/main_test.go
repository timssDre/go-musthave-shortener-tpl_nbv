package main

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_shortenURLHandler(t *testing.T) {
	type args struct {
		code        int
		contentType string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test1",
			args: args{
				code:        201,
				contentType: "text/plain",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := gin.Default()
			r.POST("/", shortenURLHandler)

			request := httptest.NewRequest(http.MethodPost, "/", nil)
			w := httptest.NewRecorder()

			r.ServeHTTP(w, request)

			res := w.Result()
			defer res.Body.Close()
			assert.Equal(t, tt.args.code, res.StatusCode)
			assert.Equal(t, tt.args.contentType, res.Header.Get("Content-Type"))
		})
	}
}

func Test_redirectToOriginalURLHandler(t *testing.T) {
	type argsGet struct {
		code     int
		location string
	}
	testsGET := []struct {
		name    string
		argsGet argsGet
	}{
		{
			name: "test1",
			argsGet: argsGet{
				code:     307,
				location: "https://practicum.yandex.ru/",
			},
		},
	}

	urlMap["asd"] = "https://practicum.yandex.ru/"

	for _, tt := range testsGET {
		t.Run(tt.name, func(t *testing.T) {
			r := gin.Default()
			r.GET("/:id", redirectToOriginalURLHandler)

			request := httptest.NewRequest(http.MethodGet, "/asd", nil)
			w := httptest.NewRecorder()
			r.ServeHTTP(w, request)
			res := w.Result()
			defer res.Body.Close()
			assert.Equal(t, tt.argsGet.code, res.StatusCode)
			assert.Equal(t, tt.argsGet.location, res.Header.Get("location"))
		})
	}
}
