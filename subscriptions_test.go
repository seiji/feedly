package feedly

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubscriptionsGet(t *testing.T) {
	expected := Subscriptions{
		{
			Categories:  []Category{},
			ContentType: "",
			IconUrl:     "",
			ID:          "feed/https://example.com.com/rss/feed/",
			Partial:     false,
			Subscribers: 0,
			Title:       "",
			Topics:      []string{},
			Updated:     0,
			Velocity:    0,
			VisualUrl:   "",
			Website:     "",
		},
	}
	client := NewTestClient(expected)
	api := NewAPI(client)
	ctx := context.Background()
	actual, err := api.SubscriptionsGet(ctx)
	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}
