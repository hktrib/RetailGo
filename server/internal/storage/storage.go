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



// // Download  retrieves a file object, if it exists, otherwise return file response
// func (f *file) Download(filePath string) ([]byte, error) {
// 	reqURL := fmt.Sprintf("%s/%s/object/authenticated/%s/%s", f.storage.client.BaseURL, StorageEndpoint, f.BucketId, filePath)
// 	req, err := http.NewRequest(http.MethodGet, reqURL, nil)
// 	if err != nil {
// 		panic(err)
// 	}

// 	injectAuthHeader(req, f.storage.client.)

// 	client := &http.Client{}
// 	res, err := client.Do(req)
// 	if err != nil {
// 		panic(err)
// 	}

// 	body, err := io.ReadAll(res.Body)
// 	if err != nil {
// 		panic(err)
// 	}

// 	// when not success, supabase will return json insted of file
// 	if res.StatusCode != 200 {
// 		var resErr *FileErrorResponse
// 		if err := json.Unmarshal(body, &resErr); err != nil {
// 			panic(err)
// 		}

// 		if resErr.Status == "404" {
// 			return nil, ErrNotFound
// 		}

// 		return nil, resErr
// 	}

// 	return body, nil
// }
