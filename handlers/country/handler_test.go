package country

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHandler_ListAll(t *testing.T) {
	token := os.Getenv("MS_TOKEN")
	handler := NewHandler(token)
	countries, err := handler.ListAll(context.Background())
	require.NoError(t, err)
	assert.NotEmpty(t, countries)
}

func TestHandler_Search(t *testing.T) {
	token := os.Getenv("MS_TOKEN")
	handler := NewHandler(token)
	countries, err := handler.Search(context.Background(), "канада")
	require.NoError(t, err)
	assert.Equal(t, 1, len(countries))
}
