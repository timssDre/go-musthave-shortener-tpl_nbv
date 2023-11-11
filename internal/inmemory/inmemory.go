package inmemory

func New(BaseURL string) *Storage {
	return &Storage{
		BaseURL: BaseURL,
		UrlMap:  make(map[string]string),
	}
}

type Storage struct {
	BaseURL string
	UrlMap  map[string]string
}
