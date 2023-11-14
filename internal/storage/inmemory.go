package storage

type Storage struct {
	URLs map[string]string
}

func New() *Storage {
	return &Storage{
		URLs: make(map[string]string),
	}
}

func (s *Storage) Get(key string) (string, bool) {
	value, exists := s.URLs[key]
	return value, exists
}

func (s *Storage) Set(key string, value string) {
	s.URLs[key] = value
}
