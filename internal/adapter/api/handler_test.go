package api

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
		name    string
		Storage Storage
		args    args
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
			tt.Storage.urlMap = make(map[string]string)
			tt.Storage.BaseURL = "http://localhost:8081"

			r := gin.Default()

			r.POST("/", tt.Storage.ShortenURLHandler)

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
		testURL  string
		location string
	}
	testsGET := []struct {
		name    string
		Storage Storage
		argsGet argsGet
	}{
		{
			name: "test1",
			argsGet: argsGet{
				code:     307,
				testURL:  "ads",
				location: "https://practicum.yandex.ru/",
			},
		},
	}

	for _, tt := range testsGET {
		t.Run(tt.name, func(t *testing.T) {
			tt.Storage.BaseURL = "http://localhost:8081"
			tt.Storage.urlMap = make(map[string]string)
			tt.Storage.urlMap[tt.argsGet.testURL] = tt.argsGet.location

			r := gin.Default()
			r.GET("/:id", tt.Storage.RedirectToOriginalURLHandler)

			request := httptest.NewRequest(http.MethodGet, "/ads", nil)
			w := httptest.NewRecorder()
			r.ServeHTTP(w, request)
			res := w.Result()
			defer res.Body.Close()
			assert.Equal(t, tt.argsGet.code, res.StatusCode)
			assert.Equal(t, tt.argsGet.location, res.Header.Get("location"))
		})
	}
}