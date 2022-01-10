package feedly

import (
	"context"
	"encoding/json"
	"net/http"
)

type apiMarkers struct {
	api *apiV3
}

type Marker struct {
	UnreadCounts []UnreadCount `json:"unreadcounts"`
	Updated      int64         `json:"updated"`
}

func (a Marker) String() string {
	e, err := json.Marshal(a)
	if err != nil {
		panic(err)
	}
	return string(e)
}

type UnreadCount struct {
	Count   int    `json:"count"`
	ID      string `json:"id"`
	Updated int64  `json:"updated"`
}

type MarkersReads struct {
	Entries []string           `json:"entries"`
	Feeds   []MarkersReadsFeed `json:"feeds"`
	Updated int64              `json:"updated"`
}

func (a MarkersReads) String() string {
	e, err := json.Marshal(a)
	if err != nil {
		panic(err)
	}
	return string(e)
}

type MarkersReadsFeed struct {
	ID   string `json:"id"`
	AsOf int64  `json:"asOf"`
}

type MarkersReadsOptions struct {
	NewerThan int64 `url:"newerThan,omitempty"`
}

func (a *apiMarkers) MarkersCounts(ctx context.Context) (marker *Marker, err error) {
	var req *http.Request
	if req, err = a.api.NewRequest(ctx, "GET", "markers/counts", nil); err != nil {
		return
	}
	marker = new(Marker)
	if _, err = a.api.Do(req, marker); err != nil {
		return
	}
	return
}

func (a *apiMarkers) MarkersReads(ctx context.Context, opt *MarkersReadsOptions) (
	markersReads *MarkersReads,
	err error,
) {
	var req *http.Request
	rel := "markers/reads"
	if rel, err = addOptions(rel, opt); err != nil {
		return
	}
	if req, err = a.api.NewRequest(ctx, "GET", rel, nil); err != nil {
		return
	}
	markersReads = new(MarkersReads)
	if _, err = a.api.Do(req, markersReads); err != nil {
		return
	}
	return
}
