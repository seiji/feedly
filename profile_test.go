package feedly

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProfileGet(t *testing.T) {
	test := Profile{
		Email: "test@example.com",
	}
	client := NewTestClient(func(req *http.Request) *http.Response {
		b, _ := json.Marshal(test)
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(b)),
			Header:     make(http.Header),
		}
	})
	api := NewAPI(client)
	got, err := api.ProfileGet(nil)
	assert.Nil(t, err)
	assert.Equal(t, got, &test)
}
