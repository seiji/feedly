package feedly

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
)

type apiSubscriptions struct {
	api *apiV3
}

type Subscription struct {
	Categories  []Category `json:"categories"`
	ContentType string     `json:"contentType"`
	IconURL     string     `json:"iconUrl"`
	ID          string     `json:"id"`
	Partial     bool       `json:"partial"`
	Subscribers float64    `json:"subscribers"`
	Title       string     `json:"title"`
	Topics      []string   `json:"topics"`
	Updated     float64    `json:"updated"`
	Velocity    float64    `json:"velocity"`
	VisualURL   string     `json:"visualUrl"`
	Website     string     `json:"website"`
}

func (a Subscription) String() string {
	e, err := json.Marshal(a)
	if err != nil {
		panic(err)
	}
	return string(e)
}

type Subscriptions []Subscription

func (a Subscriptions) String() string {
	s := make([]string, len(a))
	for i, v := range a {
		s[i] = v.String()
	}
	return "[" + strings.Join(s, ",") + "]"
}

func (a *apiSubscriptions) SubscriptionsGet(ctx context.Context) (subscriptions Subscriptions, err error) {
	var req *http.Request
	if req, err = a.api.NewRequest(ctx, "GET", "subscriptions", nil); err != nil {
		return nil, err
	}

	if _, err = a.api.Do(req, &subscriptions); err != nil {
		return nil, err
	}
	return subscriptions, nil
}
