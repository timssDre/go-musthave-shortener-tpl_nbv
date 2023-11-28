package storage

import "encoding/json"

func (s *Storage) WriteEvent(event *ShortCollector) error {
	data, err := json.Marshal(&event)
	if err != nil {
		return err
	}

	// записываем событие в буфер
	if _, err := s.writer.Write(data); err != nil {
		return err
	}

	// добавляем перенос строки
	if err := s.writer.WriteByte('\n'); err != nil {
		return err
	}

	// записываем буфер в файл
	return s.writer.Flush()
}
