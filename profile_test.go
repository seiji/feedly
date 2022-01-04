package feedly

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProfileGet(t *testing.T) {
	expected := Profile{
		Email: "test@example.com",
	}
	client := NewTestClient(func(req *http.Request) *http.Response {
		b, _ := json.Marshal(expected)
		return &http.Response{
			Status:           "",
			StatusCode:       http.StatusOK,
			Proto:            "",
			ProtoMajor:       0,
			ProtoMinor:       0,
			Header:           make(http.Header),
			Body:             ioutil.NopCloser(bytes.NewBuffer(b)),
			ContentLength:    0,
			TransferEncoding: []string{},
			Close:            false,
			Uncompressed:     false,
			Trailer:          map[string][]string{},
			Request:          &http.Request{},
			TLS:              &tls.ConnectionState{},
		}
	})
	api := NewAPI(client)
	ctx := context.Background()
	actual, err := api.ProfileGet(ctx)
	assert.Nil(t, err)
	assert.Equal(t, &expected, actual)
}
