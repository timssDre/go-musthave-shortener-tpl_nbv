package api

func New(BaseURL string) *Storage {
	return &Storage{
		BaseURL: BaseURL,
		urlMap:  make(map[string]string),
	}
}

type Storage struct {
	BaseURL string
	urlMap  map[string]string
}
