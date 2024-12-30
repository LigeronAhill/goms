package test

import (
	"context"
	"os"
	"testing"

	"github.com/LigeronAhill/goms"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMain(t *testing.T) {
	ctx := context.Background()
	token := os.Getenv("MS_TOKEN")
	if len(token) == 0 {
		t.Fatal("no token")
	}
	client := goms.New(token)
	countries, err := client.CountryHandler.ListAll(ctx)
	require.NoError(t, err)
	assert.NotEmpty(t, countries)
}
