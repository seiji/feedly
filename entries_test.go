package feedly

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEntriesGet(t *testing.T) {
	id := "Xne8uW/IUiZhV1EuO2ZMzIrc2Ak6NlhGjboZ+Yk0rJ8=_1523699cbb3:2aa0463:e47a7aef"
	expected := Entries{
		{
			Alternate:      []Alternate{},
			Author:         "",
			Categories:     []Category{},
			Content:        &Content{},
			Crawled:        0,
			Engagement:     0,
			EngagementRate: 0,
			Fingerprint:    "",
			ID:             id,
			Keywords:       []string{},
			Origin:         &Origin{},
			OriginID:       "",
			Published:      0,
			Title:          "",
			Unread:         false,
			Updated:        0,
			Visual:         &Visual{},
		},
	}
	client := NewTestClient(expected)
	api := NewAPI(client)
	ctx := context.Background()
	actual, err := api.EntriesGet(ctx, id)
	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}
