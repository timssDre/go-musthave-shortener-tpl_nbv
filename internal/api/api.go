package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/timssDre/go-musthave-shortener-tpl_nbv.git/internal/services"
	"github.com/timssDre/go-musthave-shortener-tpl_nbv.git/internal/storage"
)

type RestAPI struct {
	StructService *services.ShortenerService
}

func StartRestAPI(ServerAddr, BaseURL string, storage *storage.Storage) error {
	storageShortener := services.NewShortenerService(BaseURL, storage)
	api := &RestAPI{
		StructService: storageShortener,
	}

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	api.setRoutes(r)

	err := r.Run(ServerAddr)
	if err != nil {
		fmt.Println("failed to start the browser")
		return err
	}

	return nil
}
