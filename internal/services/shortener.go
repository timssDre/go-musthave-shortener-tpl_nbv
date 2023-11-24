package services

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/timssDre/go-musthave-shortener-tpl_nbv.git/internal/storage"
	"log"
)

type ShortenerService struct {
	BaseURL string
	Storage *storage.Storage
}

func NewShortenerService(BaseURL string, storage *storage.Storage) *ShortenerService {
	s := &ShortenerService{
		BaseURL: BaseURL,
		Storage: storage,
	}
	return s
}

func (s *ShortenerService) GetShortURL(originalURL string) (string, error) {
	shortID := randSeq()
	err := s.Storage.Set(shortID, originalURL)
	if err != nil {
		log.Fatal("failed to record event to file")
		//return "", err
	}
	shortURL := fmt.Sprintf("%s/%s", s.BaseURL, shortID)
	return shortURL, err
}

func randSeq() string {
	newUUID := uuid.New()
	return newUUID.String()
}

func (s *ShortenerService) GetOriginalURL(shortID string) (string, bool) {
	return s.Storage.Get(shortID)
}
