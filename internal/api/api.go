package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/timssDre/go-musthave-shortener-tpl_nbv.git/internal/storage"
)

type Api struct {
	BaseURL string
	storage *storage.Storage
}

func StartService(ServerAddr, BaseURL string, storage *storage.Storage) error {
	api := &Api{
		BaseURL: BaseURL,
		storage: storage,
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
