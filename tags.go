package feedly

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
)

type APITags struct {
	api *apiV3
}

type Tag struct {
	ID          string `json:"id"`
	Label       string `json:"label"`
	Description string `json:"description"`
}

func (a Tag) String() string {
	e, err := json.Marshal(a)
	if err != nil {
		panic(err)
	}
	return string(e)
}

type Tags []Tag

func (a Tags) String() string {
	s := make([]string, len(a))
	for i, v := range a {
		s[i] = v.String()
	}
	return "[" + strings.Join(s, ",") + "]"
}

func (a *APITags) TagsList(ctx context.Context) (tags Tags, err error) {
	var req *http.Request
	if req, err = a.api.NewRequest("GET", "tags", nil); err != nil {
		return nil, err
	}
	if _, err = a.api.Do(req, &tags); err != nil {
		return nil, err
	}
	return tags, nil
}
