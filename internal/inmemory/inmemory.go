package inmemory

func New(BaseURL string) *Storage {
	return &Storage{
		BaseURL: BaseURL,
		URLMap:  make(map[string]string),
	}
}

type Storage struct {
	BaseURL string
	URLMap  map[string]string
}
