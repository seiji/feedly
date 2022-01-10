package feedly

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

const priorityID = "user/f70425c6-a169-4483-8450-f4ded012354b/priority/c4b6b721-ee19-4185-b4b2-2ec990f66040"

func TestPrioritiesList(t *testing.T) {
	expected := Priorities{{
		Active:                   false,
		ActiveUntil:              0,
		Created:                  0,
		ID:                       priorityID,
		Label:                    "",
		LastEntryMatch:           0,
		LastLikeBoardRequestDate: 0,
		LastStatsRefreshDate:     0,
		LastTrainingDate:         0,
		LastUpdated:              0,
		LatestRefreshErrors:      []PriorityRefreshError{},
		Layers:                   []Filter{},
		NextRun:                  0,
		NumEntries:               0,
		NumEntriesMatching:       0,
		NumEntriesProcessed:      0,
		StreamIDs:                []string{},
		TrainingScore:            0,
		TrainingStatus:           "",
	}}
	client := NewTestClient(expected)
	api := NewAPI(client)
	ctx := context.Background()
	actual, err := api.PrioritiesList(ctx, false)
	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}

func TestPrioritiesPut(t *testing.T) {
	expected := Priorities{{
		Active:                   false,
		ActiveUntil:              0,
		Created:                  0,
		ID:                       priorityID,
		Label:                    "test",
		LastEntryMatch:           0,
		LastLikeBoardRequestDate: 0,
		LastStatsRefreshDate:     0,
		LastTrainingDate:         0,
		LastUpdated:              0,
		LatestRefreshErrors:      []PriorityRefreshError{},
		Layers:                   []Filter{},
		NextRun:                  0,
		NumEntries:               0,
		NumEntriesMatching:       0,
		NumEntriesProcessed:      0,
		StreamIDs:                []string{},
		TrainingScore:            0,
		TrainingStatus:           "",
	}}
	client := NewTestClient(expected)
	api := NewAPI(client)
	ctx := context.Background()
	actual, err := api.PrioritiesPut(ctx, PriorityCreate{
		Active:      true,
		ActiveUntil: 0,
		Label:       "test",
		Layers:      []Filter{},
		StreamIDs:   []string{},
	})
	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}

func TestPrioritiesUpdate(t *testing.T) {
	expected := Priorities{{
		Active:                   false,
		ActiveUntil:              0,
		Created:                  0,
		ID:                       priorityID,
		Label:                    "test",
		LastEntryMatch:           0,
		LastLikeBoardRequestDate: 0,
		LastStatsRefreshDate:     0,
		LastTrainingDate:         0,
		LastUpdated:              0,
		LatestRefreshErrors:      []PriorityRefreshError{},
		Layers:                   []Filter{},
		NextRun:                  0,
		NumEntries:               0,
		NumEntriesMatching:       0,
		NumEntriesProcessed:      0,
		StreamIDs:                []string{},
		TrainingScore:            0,
		TrainingStatus:           "",
	}}
	client := NewTestClient(expected)
	api := NewAPI(client)
	ctx := context.Background()
	actual, err := api.PrioritiesUpdate(ctx, PriorityUpdate{
		ID:          priorityID,
		Active:      true,
		ActiveUntil: 0,
		Label:       "test",
		Layers:      []Filter{},
		StreamIDs:   []string{},
	})
	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}

func TestPrioritiesDelete(t *testing.T) {
	client := NewTestClient(struct{}{})
	api := NewAPI(client)
	ctx := context.Background()
	err := api.PrioritiesDelete(ctx, priorityID, false)
	assert.Nil(t, err)
}
