package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/timssDre/go-musthave-shortener-tpl_nbv.git/internal/dump"
	"github.com/timssDre/go-musthave-shortener-tpl_nbv.git/internal/logger"
	"github.com/timssDre/go-musthave-shortener-tpl_nbv.git/internal/middleware"
	"github.com/timssDre/go-musthave-shortener-tpl_nbv.git/internal/services"
	"github.com/timssDre/go-musthave-shortener-tpl_nbv.git/internal/storage"
	"go.uber.org/zap"
	"log"
)

type RestAPI struct {
	StructService *services.ShortenerService
	StructDump    *dump.Memory
}

func StartRestAPI(ServerAddr, BaseURL string, LogLevel string, FilePath string, storage *storage.Storage) error {
	if err := logger.Initialize(LogLevel); err != nil {
		return err
	}
	logger.Log.Info("Running server", zap.String("address", ServerAddr))

	storageShortener := services.NewShortenerService(BaseURL, storage)
	d := dump.NewMemory()
	err := d.FillFromStorage(FilePath, storageShortener)
	if err != nil {
		log.Fatal(err)
	}

	api := &RestAPI{
		StructService: storageShortener,
		StructDump:    d,
	}

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.Use(middleware.LoggerMiddleware(logger.Log), gin.Recovery())
	r.Use(middleware.CompressMiddleware(), gin.Recovery())

	api.setRoutes(r)

	err = r.Run(ServerAddr)
	if err != nil {
		fmt.Println("failed to start the browser")
		return err
	}
	fmt.Println("nbv")
	return nil
}
