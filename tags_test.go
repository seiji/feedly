package feedly

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTagsList(t *testing.T) {
	expected := Tags{
		{
			ID:          "user/xxx/tag/global.saved",
			Label:       "",
			Description: "",
		},
	}
	client := NewTestClient(expected)
	api := NewAPI(client)
	ctx := context.Background()
	actual, err := api.TagsList(ctx)
	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}
