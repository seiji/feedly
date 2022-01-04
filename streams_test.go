package feedly

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStreamsContents(t *testing.T) {
	id := "feed/https://www.theverge.com/rss/full.xml"
	expected := StreamContents{
		Continuation: "",
		ID:           id,
		Items:        []Entry{},
		Updated:      0,
	}
	client := NewTestClient(expected)
	api := NewAPI(client)
	ctx := context.Background()
	actual, err := api.StreamsContents(ctx, id, nil)
	assert.Nil(t, err)
	assert.Equal(t, &expected, actual)
}

func TestStreamsIDs(t *testing.T) {
	id := "feed/https://www.theverge.com/rss/full.xml"
	expected := StreamIDs{
		IDs:          []string{id},
		Continuation: "",
	}
	client := NewTestClient(expected)
	api := NewAPI(client)
	ctx := context.Background()
	actual, err := api.StreamsIDs(ctx, id, nil)
	assert.Nil(t, err)
	assert.Equal(t, &expected, actual)
}
