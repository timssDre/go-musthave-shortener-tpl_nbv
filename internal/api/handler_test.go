package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/timssDre/go-musthave-shortener-tpl_nbv.git/internal/services"
	"github.com/timssDre/go-musthave-shortener-tpl_nbv.git/internal/storage"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_shortenURLHandler(t *testing.T) {
	type args struct {
		code        int
		contentType string
	}
	tests := []struct {
		name    string
		Storage RestAPI
		args    args
		body    string
	}{
		{
			name: "test1",
			Storage: RestAPI{
				StructService: &services.ShortenerService{
					Storage: &storage.Storage{},
				},
			},
			args: args{
				code:        201,
				contentType: "text/plain",
			},
			body: "https://practicum.yandex.ru/",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.Storage.StructService.Storage.URLs = make(map[string]string)
			tt.Storage.StructService.BaseURL = "http://localhost:8081"

			r := gin.Default()

			r.POST("/", tt.Storage.ShortenURLHandler)

			request := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(tt.body))
			w := httptest.NewRecorder()

			r.ServeHTTP(w, request)

			res := w.Result()
			defer res.Body.Close()

			assert.Equal(t, tt.args.code, res.StatusCode)
			assert.Equal(t, tt.args.contentType, res.Header.Get("Content-Type"))
		})
	}
}

func Test_shortenURLHandlerURL(t *testing.T) {
	type args struct {
		code        int
		contentType string
	}
	type reqBody struct {
		PerformanceURL string `json:"url"`
	}
	tests := []struct {
		name    string
		Storage RestAPI
		args    args
		body    reqBody
	}{
		{
			name: "test1",
			Storage: RestAPI{
				StructService: &services.ShortenerService{
					Storage: &storage.Storage{},
				},
			},
			args: args{
				code:        201,
				contentType: "application/json",
			},
			body: reqBody{
				"https://practicum.yandex.ru",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.Storage.StructService.Storage.URLs = make(map[string]string)
			tt.Storage.StructService.BaseURL = "http://localhost:8081"

			r := gin.Default()

			r.POST("/api/shorten", tt.Storage.ShortenURLHandlerJSON)
			jsonBody, err := json.Marshal(tt.body)
			if err != nil {
				t.Fatal(err)
			}
			request := httptest.NewRequest(http.MethodPost, "/api/shorten", strings.NewReader(string(jsonBody)))
			request.Header.Set("Content-Type", "application/json")
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
		Storage RestAPI
		argsGet argsGet
	}{
		{
			name: "test1",
			Storage: RestAPI{
				StructService: &services.ShortenerService{
					Storage: &storage.Storage{},
				},
			},
			argsGet: argsGet{
				code:     307,
				testURL:  "ads",
				location: "https://practicum.yandex.ru/",
			},
		},
	}

	for _, tt := range testsGET {
		t.Run(tt.name, func(t *testing.T) {
			tt.Storage.StructService.BaseURL = "http://localhost:8081"
			tt.Storage.StructService.Storage.URLs = make(map[string]string)
			tt.Storage.StructService.Storage.Set(tt.argsGet.testURL, tt.argsGet.location)

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
