package storage

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
)

type Storage struct {
	URLs    map[string]string
	file    *os.File
	writer  *bufio.Writer
	maxUUID int
}

type ShortCollector struct {
	NumberUUID  string `json:"uuid"`
	ShortURL    string `json:"short_url"`
	OriginalURL string `json:"original_url"`
}

func NewStorage() *Storage {
	return &Storage{
		URLs: make(map[string]string),
	}
}

func (s *Storage) FillFromStorage(filePath string) error {
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	newDecoder := json.NewDecoder(file)
	maxUUID := 0
	for {
		var event ShortCollector
		if err := newDecoder.Decode(&event); err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Println("error decode JSON:", err)
				break
			}
		}
		maxUUID += 1
		s.URLs[event.OriginalURL] = event.ShortURL
	}
	s.file = file
	s.maxUUID = maxUUID
	s.writer = bufio.NewWriter(file)
	return nil
}

func (s *Storage) Set(key string, value string) error {
	s.URLs[key] = value
	if s.file == nil {
		return nil
	}
	s.maxUUID += 1
	event := ShortCollector{
		strconv.Itoa(s.maxUUID),
		key,
		value,
	}
	err := s.WriteEvent(&event)

	return err
}

func (s *Storage) Get(key string) (string, bool) {
	value, exists := s.URLs[key]
	return value, exists
}
