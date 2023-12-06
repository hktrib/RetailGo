package s3

import (
	"fmt"
	"io"
	"net/http"

	"github.com/hktrib/RetailGo/internal/ent"
)

const (
	StorageEndpoint = "storage/v1"
)

type Storage struct {
	BaseURL string
	// apiKey can be a client API key or a service key
	apiKey     string
	HTTPClient *http.Client
	DB         *ent.Client
}

func NewStorage(baseURL string, apiKey string, client *ent.Client) *Storage {
	return &Storage{
		BaseURL: baseURL,
		apiKey: apiKey,
		HTTPClient: &http.Client{},
		DB: client,
	}
}

func (s *Storage) NewRequestWithAuthHeader(method string, url string, body io.Reader) (*http.Request, error) {
	reqUrl := fmt.Sprintf("%s/%s/%s", s.BaseURL, StorageEndpoint, url)
	req, err := http.NewRequest(http.MethodDelete, reqUrl, body)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func (s *Storage) injectAuthHeader(req *http.Request) {
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", s.apiKey))
}
