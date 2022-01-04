package feedly

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
)

type APIFeeds struct {
	api *apiV3
}

type Feed struct {
	ID          string   `json:"id"`
	Title       string   `json:"title"`
	Topics      []string `json:"topics"`
	Language    string   `json:"language"`
	Website     string   `json:"website"`
	Velocity    float64  `json:"velocity"`
	Featured    bool     `json:"featured"`
	Sponsored   bool     `json:"sponsored"`
	Curated     bool     `json:"curated"`
	Subscribers int64    `json:"subscribers"`
	State       string   `json:"alive"`
}

func (a Feed) String() string {
	e, err := json.Marshal(a)
	if err != nil {
		panic(err)
	}
	return string(e)
}

type Feeds []Feed

func (a Feeds) String() string {
	s := make([]string, len(a))
	for i, v := range a {
		s[i] = v.String()
	}
	return "[" + strings.Join(s, ",") + "]"
}

func (a *APIFeeds) FeedsGet(ctx context.Context, feedID string) (feed *Feed, err error) {
	var req *http.Request
	rel := "feeds/" + url.QueryEscape(feedID)
	if req, err = a.api.NewRequest("GET", rel, nil); err != nil {
		return nil, err
	}
	feed = new(Feed)
	if _, err := a.api.Do(req, feed); err != nil {
		return nil, err
	}
	return feed, nil
}

func (a *APIFeeds) FeedsMGet(ctx context.Context, feedIDs []string) (feeds Feeds, err error) {
	var req *http.Request
	rel := "feeds/.mget"
	if req, err = a.api.NewRequest("POST", rel, feedIDs); err != nil {
		return nil, err
	}
	if _, err = a.api.Do(req, &feeds); err != nil {
		return nil, err
	}
	return feeds, nil
}
