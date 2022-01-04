package feedly

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
)

type APIEntries struct {
	api *apiV3
}

type Alternate struct {
	Href string `json:"href"`
	Type string `json:"type"`
}

type Content struct {
	Content   string `json:"content"`
	Direction string `json:"direction"`
}

type Origin struct {
	HtmlUrl  string `json:"htmlUrl"`
	StreamId string `json:"streamId"`
	Title    string `json:"title"`
}

type Visual struct {
	ContentType string  `json:"contentType"`
	Height      float64 `json:"height"`
	Processor   string  `json:"processor"`
	Url         string  `json:"url"`
	Width       float64 `json:"width"`
}

type Entry struct {
	Alternate      []Alternate `json:"alternate"`
	Author         string      `json:"author"`
	Categories     []Category  `json:"categories"`
	Content        *Content    `json:"content"`
	Crawled        int64       `json:"crawled"`
	Engagement     int64       `json:"engagement"`
	EngagementRate float64     `json:"engagementRate"`
	Fingerprint    string      `json:"fingerprint"`
	ID             string      `json:"id"`
	Keywords       []string    `json:"keywords"`
	Origin         *Origin     `json:"origin"`
	OriginId       string      `json:"originId"`
	Published      int64       `json:"published"`
	Title          string      `json:"title"`
	Unread         bool        `json:"unread"`
	Updated        int64       `json:"updated"`
	Visual         *Visual     `json:"visual"`
}

func (a Entry) String() string {
	e, err := json.Marshal(a)
	if err != nil {
		panic(err)
	}
	return string(e)
}

type Entries []Entry

func (a Entries) String() string {
	s := make([]string, len(a))
	for i, v := range a {
		s[i] = v.String()
	}
	return "[" + strings.Join(s, ",") + "]"
}

func (a *APIEntries) EntriesGet(ctx context.Context, entryID string) (entries Entries, err error) {
	var req *http.Request
	if req, err = a.api.NewRequest("GET", "entries/"+entryID, nil); err != nil {
		return nil, err
	}
	if _, err := a.api.Do(req, &entries); err != nil {
		return nil, err
	}
	return entries, nil
}

func (a *APIEntries) MGet(entryIds []string) ([]Entry, *Response, error) {
	rel := "entries/.mget"
	req, err := a.api.NewRequest("POST", rel, entryIds)
	if err != nil {
		return nil, nil, err
	}

	entries := new([]Entry)

	res, err := a.api.Do(req, entries)
	if err != nil {
		return nil, res, err
	}

	return *entries, res, nil
}
