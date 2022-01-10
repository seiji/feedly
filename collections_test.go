package feedly

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

const collectionID = "user/af190c49-0ac8-4f08-9f83-805f1a3bc142/category/c805fcbf-3acf-4302-a97e-d82f9d7c897f"

func TestCollectionsCreate(t *testing.T) {
	expected := Collections{{
		Customizable: false,
		Description:  "",
		Enterprise:   false,
		Feeds:        []Feed{{ID: "feed/http://feeds.feedburner.com/design-milk", Title: ""}},
		ID:           collectionID,
		Label:        "",
		NumFeeds:     1,
	}}
	client := NewTestClient(expected)
	api := NewAPI(client)
	ctx := context.Background()
	collections, err := api.CollectionsCreate(ctx, &CollectionCreate{
		Description: "",
		Feeds:       []CollectionFeedCreate{{ID: "feed/http://feeds.feedburner.com/design-milk", Title: ""}},
		ID:          collectionID,
		Label:       "Feedly",
	})
	assert.Nil(t, err)
	assert.Equal(t, 1, len(collections))
}

func TestCollectionsPut(t *testing.T) {
	expected := Feeds{{
		ContentType:         "",
		Description:         "",
		EstimatedEngagement: 0,
		FeedID:              "feed/http://feeds.feedburner.com/design-milk",
		IconURL:             "",
		ID:                  "feed/http://feeds.feedburner.com/design-milk",
		Language:            "",
		Partial:             false,
		Subscribers:         0,
		Title:               "",
		Topics:              []string{},
		Updated:             0,
		Velocity:            0,
		VisualURL:           "",
		Website:             "",
	}}
	client := NewTestClient(expected)
	api := NewAPI(client)
	ctx := context.Background()
	feeds, err := api.CollectionsFeedsPut(ctx, collectionID, CollectionFeedCreate{
		ID:    "feed/http://feeds.feedburner.com/design-milk",
		Title: "",
	})
	assert.Nil(t, err)
	assert.Equal(t, 1, len(feeds))
}

func TestCollectionsMPut(t *testing.T) {
	expected := Feeds{{
		ContentType:         "",
		Description:         "",
		EstimatedEngagement: 0,
		FeedID:              "feed/http://feeds.feedburner.com/design-milk",
		IconURL:             "",
		ID:                  "feed/http://feeds.feedburner.com/design-milk",
		Language:            "",
		Partial:             false,
		Subscribers:         0,
		Title:               "",
		Topics:              []string{},
		Updated:             0,
		Velocity:            0,
		VisualURL:           "",
		Website:             "",
	}}
	client := NewTestClient(expected)
	api := NewAPI(client)
	ctx := context.Background()
	feeds, err := api.CollectionsFeedsMPut(ctx, collectionID, CollectionFeedCreates{
		{
			ID:    "feed/http://feeds.feedburner.com/design-milk",
			Title: "",
		},
	})
	assert.Nil(t, err)
	assert.Equal(t, 1, len(feeds))
}

func TestCollectionsDelete(t *testing.T) {
	client := NewTestClient(struct{}{})
	api := NewAPI(client)
	ctx := context.Background()
	err := api.CollectionsFeedsDelete(ctx, collectionID, "feed/http://feeds.feedburner.com/design-milk")
	assert.Nil(t, err)
}

func TestCollectionsMDelete(t *testing.T) {
	client := NewTestClient(struct{}{})
	api := NewAPI(client)
	ctx := context.Background()
	err := api.CollectionsFeedsMDelete(ctx, collectionID, CollectionFeedDeletes{
		{ID: "feed/http://feeds.feedburner.com/design-milk"},
	})
	assert.Nil(t, err)
}

func TestCollectionsGet(t *testing.T) {
	expected := Collections{{
		Customizable: false,
		Description:  "",
		Enterprise:   false,
		Feeds:        []Feed{},
		ID:           collectionID,
		Label:        "",
		NumFeeds:     0,
	}}
	client := NewTestClient(expected)
	api := NewAPI(client)
	ctx := context.Background()
	actual, err := api.CollectionsGet(ctx, collectionID)
	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}

func TestCollectionsList(t *testing.T) {
	expected := Collections{{
		Customizable: false,
		Enterprise:   false,
		Feeds:        []Feed{},
		ID:           collectionID,
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
