package gomystorage

import (
	"net/http"
	"net/url"
	"path"
)

const host = "api.moysklad.ru"

type ApiClient struct {
	baseUrl url.URL
	client  http.Client
	token   string
}

// New creates a new ApiClient instance with the provided token.
//
// token is the authentication token to be used for API requests.
// Returns a pointer to the newly created ApiClient instance.
func New(token string) *ApiClient {
	baseUrl := url.URL{
		Scheme: "https",
		Host:   host,
		Path:   path.Join("/api", "remap", "1.2"),
	}
	token = "Bearer " + token
	return &ApiClient{
		baseUrl: baseUrl,
		client:  http.Client{},
		token:   token,
	}
}
