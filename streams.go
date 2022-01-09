package feedly

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
)

type apiStreams struct {
	api *apiV3
}

type StreamContents struct {
	Continuation string  `json:"continuation"`
	ID           string  `json:"id"`
	Items        []Entry `json:"items"`
	Updated      int64   `json:"updated"`
}

func (a StreamContents) String() string {
	e, err := json.Marshal(a)
	if err != nil {
		panic(err)
	}
	return string(e)
}

type StreamIDs struct {
	IDs          []string `json:"ids"`
	Continuation string   `json:"continuation,omitempty"`
}

func (a StreamIDs) String() string {
	e, err := json.Marshal(a)
	if err != nil {
		panic(err)
	}
	return string(e)
}

type StreamOptions struct {
	Count        uint   `url:"count,omitempty"`
	Ranked       string `url:"ranked,omitempty"`
	UnreadOnly   bool   `url:"unreadOnly,omitempty"`
	NewerThan    int64  `url:"newerThan,omitempty"`
	Continuation string `url:"continuation,omitempty"`
}

func (a *apiStreams) StreamsContents(ctx context.Context, id string, opt *StreamOptions) (
	streamContents *StreamContents,
	err error,
) {
	rel := "streams/" + url.QueryEscape(id) + "/contents"
	if rel, err = addOptions(rel, opt); err != nil {
		return
	}
	var req *http.Request
	if req, err = a.api.NewRequest("GET", rel, nil); err != nil {
		return
	}
	streamContents = new(StreamContents)
	if _, err = a.api.Do(req, streamContents); err != nil {
		return
	}
	return
}

func (a *apiStreams) StreamsIDs(ctx context.Context, id string, opt *StreamOptions) (
	streamIDs *StreamIDs,
	err error,
) {
	rel := "streams/" + url.QueryEscape(id) + "/ids"
	if rel, err = addOptions(rel, opt); err != nil {
		return
	}
	var req *http.Request
	if req, err = a.api.NewRequest("GET", rel, nil); err != nil {
		return
	}
	streamIDs = new(StreamIDs)
	if _, err = a.api.Do(req, streamIDs); err != nil {
		return
	}
	return
}
