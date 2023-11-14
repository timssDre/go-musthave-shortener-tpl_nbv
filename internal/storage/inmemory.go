package storage

type Storage struct {
	URLMap map[string]string
}

func New() *Storage {
	return &Storage{
		URLMap: make(map[string]string),
	}
}

func (s *Storage) GetValueMap(key string) (string, bool) {
	value, exists := s.URLMap[key]
	return value, exists
}

func (s *Storage) SetValueMap(key string, value string) {
	s.URLMap[key] = value
}
