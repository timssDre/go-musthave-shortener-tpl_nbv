package services

import (
	"fmt"
	"github.com/timssDre/go-musthave-shortener-tpl_nbv.git/internal/storage"
	"math/rand"
)

type StructService struct {
	BaseURL string
	Storage *storage.Storage
}

func New(BaseURL string, storage *storage.Storage) *StructService {
	s := &StructService{
		BaseURL: BaseURL,
		Storage: storage,
	}
	return s
}

func (s *StructService) GetShortURL(originalURL string) string {
	shortID := randSeq(8)
	s.Storage.Set(shortID, originalURL)
	shortURL := fmt.Sprintf("%s/%s", s.BaseURL, shortID)
	return shortURL
}

func randSeq(n int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func (s *StructService) GetOriginalURL(shortID string) (string, bool) {
	return s.Storage.Get(shortID)
}
