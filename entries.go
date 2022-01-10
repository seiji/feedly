package feedly

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
)

type apiEntries struct {
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
	HTMLURL  string `json:"htmlUrl"`
	StreamID string `json:"streamId"`
	Title    string `json:"title"`
}

type Visual struct {
	ContentType string  `json:"contentType"`
	Height      float64 `json:"height"`
	Processor   string  `json:"processor"`
	URL         string  `json:"url"`
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
	OriginID       string      `json:"originId"`
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

func (a *apiEntries) EntriesGet(ctx context.Context, id string) (entries Entries, err error) {
	var req *http.Request
	if req, err = a.api.NewRequest(ctx, "GET", "entries/"+id, nil); err != nil {
		return
	}
	if _, err = a.api.Do(req, &entries); err != nil {
		return
	}
	return
}

func (a *apiEntries) MGet(ctx context.Context, ids []string) (entries Entries, err error) {
	var req *http.Request
	if req, err = a.api.NewRequest(ctx, "POST", "entries/.mget", ids); err != nil {
		return
	}
	if _, err = a.api.Do(req, &entries); err != nil {
		return
	}
	return
}
