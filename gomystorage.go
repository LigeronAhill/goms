package gomystorage

import (
	"github.com/LigeronAhill/goms/handlers/country"
)

type ApiClient struct {
	countryHandler *country.Handler
}

// New creates a new ApiClient instance with the provided token.
//
// token is the authentication token to be used for API requests.
// Returns a pointer to the newly created ApiClient instance.
func New(token string) *ApiClient {
	countryHandler := country.NewHandler(token)
	return &ApiClient{
		countryHandler,
	}
}
