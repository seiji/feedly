package feedly

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMarkersCounts(t *testing.T) {
	expected := Marker{
		UnreadCounts: []UnreadCount{},
		Updated:      0,
	}
	client := NewTestClient(expected)
	api := NewAPI(client)
	ctx := context.Background()
	actual, err := api.MarkersCounts(ctx)
	assert.Nil(t, err)
	assert.Equal(t, &expected, actual)
}

func TestMarkersReads(t *testing.T) {
	expected := MarkersReads{
		Entries: []string{},
		Feeds:   []MarkersReadsFeed{},
		Updated: 0,
	}
	client := NewTestClient(expected)
	api := NewAPI(client)
	ctx := context.Background()
	actual, err := api.MarkersReads(ctx, nil)
	assert.Nil(t, err)
	assert.Equal(t, &expected, actual)
}
