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
	ContentType         string   `json:"contentType"`
	Description         string   `json:"description"`
	EstimatedEngagement float64  `json:"estimatedEngagement"`
	FeedID              string   `json:"feedId"`
	IconURL             string   `json:"iconUrl"`
	ID                  string   `json:"id"`
	Language            string   `json:"language"`
	Partial             bool     `json:"partial"`
	Subscribers         float64  `json:"subscribers"`
	Title               string   `json:"title"`
	Topics              []string `json:"topics"`
	Updated             float64  `json:"updated"`
	Velocity            float64  `json:"velocity"`
	VisualURL           string   `json:"visualUrl"`
	Website             string   `json:"website"`
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
