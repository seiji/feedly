package feedly

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type apiPriorities struct {
	api *apiV3
}

type Priority struct {
	Active                   bool                   `json:"active"`
	ActiveUntil              float64                `json:"activeUntil"`
	Created                  float64                `json:"created"`
	ID                       string                 `json:"id"`
	Label                    string                 `json:"label"`
	LastEntryMatch           float64                `json:"lastEntryMatch"`
	LastLikeBoardRequestDate float64                `json:"lastLikeBoardRequestDate"`
	LastStatsRefreshDate     float64                `json:"lastStatsRefreshDate"`
	LastTrainingDate         float64                `json:"lastTrainingDate"`
	LastUpdated              float64                `json:"lastUpdated"`
	LatestRefreshErrors      []PriorityRefreshError `json:"latestRefreshErrors"`
	Layers                   []Filter               `json:"layers"`
	NextRun                  float64                `json:"nextRun"`
	NumEntries               float64                `json:"numEntries"`
	NumEntriesMatching       float64                `json:"numEntriesMatching"`
	NumEntriesProcessed      float64                `json:"numEntriesProcessed"`
	StreamIDs                []string               `json:"streamIds"`
	TrainingScore            float64                `json:"trainingScore"`
	TrainingStatus           string                 `json:"trainingStatus"`
}

func (a Priority) String() string {
	e, err := json.Marshal(a)
	if err != nil {
		panic(err)
	}
	return string(e)
}

type Priorities []Priority

func (a Priorities) String() string {
	s := make([]string, len(a))
	for i, v := range a {
		s[i] = v.String()
	}
	return "[" + strings.Join(s, ",") + "]"
}

type Filter struct {
	Boards        []Board      `json:"boards"`
	ComplexFilter bool         `json:"complexFilter"`
	Parts         []FilterPart `json:"parts"`
	Salience      string       `json:"salience"`
	Severities    []string     `json:"severities"`
	Type          string       `json:"type"`
}

type FilterPart struct {
	ID    string `json:"id"`
	Label string `json:"label"`
	Text  string `json:"text"`
}

type PriorityRefreshError struct {
	BoardID            string  `json:"boardId"`
	Message            string  `json:"message"`
	NumEntries         float64 `json:"numEntries"`
	NumRequiredEntries float64 `json:"numRequiredEntries"`
	Type               string  `json:"type"`
}

type PriorityCreate struct {
	Active      bool     `json:"active,omitempty"`
	ActiveUntil float64  `json:"activeUntil,omitempty"`
	Label       string   `json:"label"`
	Layers      []Filter `json:"layers"`
	StreamIDs   []string `json:"streamIds,omitempty"`
}

type PriorityCreates []PriorityCreate

type PriorityUpdate struct {
	ID          string   `json:"id"`
	Active      bool     `json:"active,omitempty"`
	ActiveUntil float64  `json:"activeUntil,omitempty"`
	Label       string   `json:"label"`
	Layers      []Filter `json:"layers"`
	StreamIDs   []string `json:"streamIds,omitempty"`
}

func (a *apiPriorities) PrioritiesList(ctx context.Context, includeDeleted bool) (priorities Priorities, err error) {
	var req *http.Request
	rel := fmt.Sprintf("priorities?includeDeleted=%s", strconv.FormatBool(includeDeleted))
	if req, err = a.api.NewRequest(ctx, "GET", rel, nil); err != nil {
		return
	}
	if _, err = a.api.Do(req, &priorities); err != nil {
		return
	}
	return
}

func (a *apiPriorities) PrioritiesPut(ctx context.Context, param PriorityCreate) (
	priorities Priorities,
	err error,
) {
	var req *http.Request
	if req, err = a.api.NewRequest(ctx, "POST", "priorities", param); err != nil {
		return
	}
	if _, err = a.api.Do(req, &priorities); err != nil {
		return
	}
	return
}

func (a *apiPriorities) PrioritiesUpdate(ctx context.Context, param PriorityUpdate) (
	priorities Priorities,
	err error,
) {
	var req *http.Request
	if req, err = a.api.NewRequest(ctx, "POST", "priorities", param); err != nil {
		return
	}
	if _, err = a.api.Do(req, &priorities); err != nil {
		return
	}
	return
}

func (a *apiPriorities) PrioritiesDelete(ctx context.Context, id string, reset bool) (err error) {
	var req *http.Request
	rel := fmt.Sprintf("priorities/%s?reset=%s", url.QueryEscape(id), strconv.FormatBool(reset))
	if req, err = a.api.NewRequest(ctx, "DELETE", rel, nil); err != nil {
		return
	}
	if _, err = a.api.Do(req, ioutil.Discard); err != nil {
		return
	}
	return
}
