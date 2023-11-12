package server

func New(BaseURL string) *Storage {
	return &Storage{
		BaseURL: BaseURL,
		URLMap:  make(map[string]string),
	}
}

func (s *Storage) GetValueMap(key string) (string, bool) {
	value, exists := s.URLMap[key]
	return value, exists
}

func (s *Storage) SetValueMap(key string, value string) {
	s.URLMap[key] = value
}

type Storage struct {
	BaseURL string
	URLMap  map[string]string
}
