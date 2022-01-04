package feedly

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFeedsGet(t *testing.T) {
	id := "feed/http://feeds.engadget.com/weblogsinc/engadget"
	expected := Feed{ID: id}
	client := NewTestClient(expected)
	api := NewAPI(client)
	ctx := context.Background()
	actual, err := api.FeedsGet(ctx, id)
	assert.Nil(t, err)
	assert.Equal(t, &expected, actual)
}

func TestFeedsMGet(t *testing.T) {
	id := "feed/http://feeds.engadget.com/weblogsinc/engadget"
	expected := Feeds{{ID: id}}
	client := NewTestClient(expected)
	api := NewAPI(client)
	ctx := context.Background()
	actual, err := api.FeedsMGet(ctx, []string{id})
	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}
