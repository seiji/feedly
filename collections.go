package feedly

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type APICollections struct {
	api *apiV3
}

type Collection struct {
	Customizable bool    `json:"customizable"`
	Description  string  `json:"description"`
	Enterprise   bool    `json:"enterprise"`
	Feeds        []Feed  `json:"feeds"`
	ID           string  `json:"id"`
	Label        string  `json:"label"`
	NumFeeds     float64 `json:"numFeeds"`
}

func (a Collection) String() string {
	e, err := json.Marshal(a)
	if err != nil {
		panic(err)
	}
	return string(e)
}

type Collections []Collection

func (a Collections) String() string {
	s := make([]string, len(a))
	for i, v := range a {
		s[i] = v.String()
	}
	return "[" + strings.Join(s, ",") + "]"
}

type CollectionCreate struct {
	Description string                 `json:"description"`
	Feeds       []CollectionCreateFeed `json:"feeds"`
	ID          string                 `json:"id"`
	Label       string                 `json:"label"`
}

type CollectionCreateFeed struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

func (a *APICollections) CollectionsCreate(ctx context.Context, param *CollectionCreate) (err error) {
	var req *http.Request
	rel := "collections"
	if req, err = a.api.NewRequest("POST", rel, param); err != nil {
		return err
	}
	if _, err = a.api.Do(req, ioutil.Discard); err != nil {
		return err
	}
	return nil
}

func (a *APICollections) CollectionsGet(ctx context.Context, collectionID string) (collections Collections, err error) {
	var req *http.Request
	rel := "collections/" + url.QueryEscape(collectionID)
	if req, err = a.api.NewRequest("GET", rel, nil); err != nil {
		return nil, err
	}
	if _, err = a.api.Do(req, &collections); err != nil {
		return nil, err
	}
	return collections, nil
}

func (a *APICollections) CollectionsList(ctx context.Context) (collections Collections, err error) {
	var req *http.Request
	if req, err = a.api.NewRequest("GET", "collections", nil); err != nil {
		return nil, err
	}

	if _, err = a.api.Do(req, &collections); err != nil {
		return nil, err
	}
	return collections, nil
}
