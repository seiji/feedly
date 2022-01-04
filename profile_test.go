package feedly

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProfileGet(t *testing.T) {
	expected := Profile{
		Email: "test@example.com",
	}
	client := NewTestClient(expected)
	api := NewAPI(client)
	ctx := context.Background()
	actual, err := api.ProfileGet(ctx)
	assert.Nil(t, err)
	assert.Equal(t, &expected, actual)
}
