package services

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/timssDre/go-musthave-shortener-tpl_nbv.git/internal/storage"
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

func (s *ShortenerService) GetShortURL(originalURL string) string {
	shortID := randSeq()
	s.Storage.Set(shortID, originalURL)
	shortURL := fmt.Sprintf("%s/%s", s.BaseURL, shortID)
	return shortURL
}

func randSeq() string {
	newUUID := uuid.New()
	return newUUID.String()
}

func (s *ShortenerService) GetOriginalURL(shortID string) (string, bool) {
	return s.Storage.Get(shortID)
}
