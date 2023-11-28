package dump

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/timssDre/go-musthave-shortener-tpl_nbv.git/internal/services"
	"io"
	"os"
	"strconv"
)

type Memory struct {
	Storage *services.ShortenerService
	file    *os.File
	writer  *bufio.Writer
	maxUUID int
}
type ShortCollector struct {
	NumberUUID  string `json:"uuid"`
	ShortURL    string `json:"short_url"`
	OriginalURL string `json:"original_url"`
}

func NewMemory() *Memory {
	return &Memory{
		maxUUID: 0,
	}
}

func (m *Memory) FillFromStorage(filePath string, storage *services.ShortenerService) error {
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
		storage.Storage.URLs[event.OriginalURL] = event.ShortURL
	}
	m.file = file
	m.maxUUID = maxUUID
	m.writer = bufio.NewWriter(file)
	return nil
}

func (m *Memory) Set(key, value string) error {
	m.maxUUID += 1
	event := ShortCollector{
		strconv.Itoa(m.maxUUID),
		key,
		value,
	}
	err := m.writeEvent(&event)
	return err
}

func (m *Memory) writeEvent(event *ShortCollector) error {
	data, err := json.Marshal(&event)
	if err != nil {
		return err
	}

	// записываем событие в буфер
	if _, err := m.writer.Write(data); err != nil {
		return err
	}

	// добавляем перенос строки
	if err := m.writer.WriteByte('\n'); err != nil {
		return err
	}

	// записываем буфер в файл
	return m.writer.Flush()
}
