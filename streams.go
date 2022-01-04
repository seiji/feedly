package feedly

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
)

type APIStreams struct {
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

func (a *APIStreams) StreamsContents(ctx context.Context, streamId string, opt *StreamOptions) (streamContents *StreamContents, err error) {
	rel := "streams/" + url.QueryEscape(streamId) + "/contents"
	if rel, err = addOptions(rel, opt); err != nil {
		return nil, err
	}
	var req *http.Request
	if req, err = a.api.NewRequest("GET", rel, nil); err != nil {
		return nil, err
	}
	streamContents = new(StreamContents)
	if _, err := a.api.Do(req, streamContents); err != nil {
		return nil, err
	}
	return streamContents, nil
}

func (a *APIStreams) StreamsIDs(ctx context.Context, streamId string, opt *StreamOptions) (streamIDs *StreamIDs, err error) {
	rel := "streams/" + url.QueryEscape(streamId) + "/ids"
	if rel, err = addOptions(rel, opt); err != nil {
		return nil, err
	}
	var req *http.Request
	if req, err = a.api.NewRequest("GET", rel, nil); err != nil {
		return nil, err
	}
	streamIDs = new(StreamIDs)
	if _, err := a.api.Do(req, streamIDs); err != nil {
		return nil, err
	}
	return streamIDs, nil
}
