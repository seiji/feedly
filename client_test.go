package feedly

import (
	"net/http"
)

type RoundTripFunc func(req *http.Request) *http.Response

func (fn RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return fn(req), nil
}

func NewTestClient(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: RoundTripFunc(fn),
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return nil
		},
		Jar:     nil,
		Timeout: 0,
	}
}
