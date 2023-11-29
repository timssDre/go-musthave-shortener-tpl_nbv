package api

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/timssDre/go-musthave-shortener-tpl_nbv.git/internal/logger"
	"github.com/timssDre/go-musthave-shortener-tpl_nbv.git/internal/middleware"
	"github.com/timssDre/go-musthave-shortener-tpl_nbv.git/internal/services"
	"github.com/timssDre/go-musthave-shortener-tpl_nbv.git/internal/storage"
	"go.uber.org/zap"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type RestAPI struct {
	StructService *services.ShortenerService
}

func StartRestAPI(ServerAddr, BaseURL string, LogLevel string, storage *storage.Storage) error {
	if err := logger.Initialize(LogLevel); err != nil {
		return err
	}
	logger.Log.Info("Running server", zap.String("address", ServerAddr))

	storageShortener := services.NewShortenerService(BaseURL, storage)
	api := &RestAPI{
		StructService: storageShortener,
	}

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.Use(middleware.LoggerMiddleware(logger.Log), gin.Recovery())
	r.Use(middleware.CompressMiddleware(), gin.Recovery())

	api.setRoutes(r)

	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	go func() {
		err := r.Run(ServerAddr)
		if err != nil {
			fmt.Println("failed to start the browser")
		}
	}()
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Ошибка при остановке сервера: %v\n", err)
	}

	return nil
}
