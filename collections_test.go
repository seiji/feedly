package feedly

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCollectionsCreate(t *testing.T) {
	id := "user/af190c49-0ac8-4f08-9f83-805f1a3bc142/category/c805fcbf-3acf-4302-a97e-d82f9d7c897f"
	expected := Collections{{
		Customizable: false,
		Description:  "",
		Enterprise:   false,
		Feeds:        []Feed{{ID: "feed/http://feeds.feedburner.com/design-milk", Title: ""}},
		ID:           id,
		Label:        "",
		NumFeeds:     1,
	}}
	client := NewTestClient(expected)
	api := NewAPI(client)
	ctx := context.Background()
	collections, err := api.CollectionsCreate(ctx, &CollectionCreate{
		Description: "",
		Feeds:       []CollectionFeedCreate{{ID: "feed/http://feeds.feedburner.com/design-milk", Title: ""}},
		ID:          id,
		Label:       "Feedly",
	})
	assert.Nil(t, err)
	assert.Equal(t, 1, len(collections))
}

func TestCollectionsDelete(t *testing.T) {
	id := "user/af190c49-0ac8-4f08-9f83-805f1a3bc142/category/c805fcbf-3acf-4302-a97e-d82f9d7c897f"
	client := NewTestClient(struct{}{})
	api := NewAPI(client)
	ctx := context.Background()
	err := api.CollectionsFeedsDelete(ctx, id, "feed/http://feeds.feedburner.com/design-milk")
	assert.Nil(t, err)
}

func TestCollectionsMDelete(t *testing.T) {
	id := "user/af190c49-0ac8-4f08-9f83-805f1a3bc142/category/c805fcbf-3acf-4302-a97e-d82f9d7c897f"
	client := NewTestClient(struct{}{})
	api := NewAPI(client)
	ctx := context.Background()
	err := api.CollectionsFeedsMDelete(ctx, id, CollectionFeedDeletes{
		{ID: "feed/http://feeds.feedburner.com/design-milk"},
	})
	assert.Nil(t, err)
}

func TestCollectionsGet(t *testing.T) {
	id := "user/af190c49-0ac8-4f08-9f83-805f1a3bc142/category/c805fcbf-3acf-4302-a97e-d82f9d7c897f"
	expected := Collections{{
		Customizable: false,
		Description:  "",
		Enterprise:   false,
		Feeds:        []Feed{},
		ID:           id,
		Label:        "",
		NumFeeds:     0,
	}}
	client := NewTestClient(expected)
	api := NewAPI(client)
	ctx := context.Background()
	actual, err := api.CollectionsGet(ctx, id)
	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}

func TestCollectionsList(t *testing.T) {
	id := "user/af190c49-0ac8-4f08-9f83-805f1a3bc142/category/c805fcbf-3acf-4302-a97e-d82f9d7c897f"
	expected := Collections{{
		Customizable: false,
		Enterprise:   false,
		Feeds:        []Feed{},
		ID:           id,
		Label:        "",
		NumFeeds:     0,
	}}
	client := NewTestClient(expected)
	api := NewAPI(client)
	ctx := context.Background()
	actual, err := api.CollectionsList(ctx)
	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}