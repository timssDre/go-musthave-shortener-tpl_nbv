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

type Event struct {
	Uuid        string `json:"uuid"`
	ShortURL    string `json:"short_url"`
	OriginalURL string `json:"original_url"`
}

func NewStorage(filePath string) (*Storage, error) {
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}
	newDecoder := json.NewDecoder(file)
	URLs := make(map[string]string)
	maxUUID := 0

	answer := &Storage{
		URLs:   URLs,
		file:   file,
		writer: bufio.NewWriter(file),
	}

	for {
		var event Event
		if err := newDecoder.Decode(&event); err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Println("error decode JSON:", err)
				break
			}
		}
		maxUUID += 1
		URLs[event.OriginalURL] = event.ShortURL
	}
	answer.maxUUID = maxUUID
	return answer, nil
}

func (s *Storage) Set(key string, value string) error {
	s.URLs[key] = value
	if s.file == nil {
		return nil
	}
	s.maxUUID += 1
	event := Event{
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
