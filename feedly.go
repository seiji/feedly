package feedly

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path"
	"reflect"
	"strconv"
	"time"

	"github.com/google/go-querystring/query"
)

const (
	version      = "v3"
	baseURLCloud = "https://cloud.feedly.com"
	// baseURLSandbox = "https://sandbox7.feedly.com"

	// headerDate      = "Date"
	headerRateCount = "X-Ratelimit-Count"
	headerRateLimit = "X-Ratelimit-Limit"
	headerRateReset = "X-RateLimit-Reset"
)

type GlobalResource int

const (
	GLOBAL_MUST GlobalResource = 1 << iota
	GLOBAL_ALL
	GLOBAL_UNCATEGORIZED
	GLOBAL_READ
	GLOBAL_SAVED
)

func (t GlobalResource) String() string {
	switch t {
	case GLOBAL_MUST:
		return "user/%s/category/global.must"
	case GLOBAL_ALL:
		return "user/%s/category/global.all"
	case GLOBAL_UNCATEGORIZED:
		return "user/%s/category/global.uncategorized"
	case GLOBAL_READ:
		return "user/%s/tag/global.read"
	case GLOBAL_SAVED:
		return "user/%s/tag/global.saved"
	}
	return ""
}

type ResourceType int

const (
	RESOURCE_FEED ResourceType = 1 << iota
	RESOURCE_CATEGORY
	RESOURCE_TAG
)

type API interface {
	ProfileGet(ctx context.Context) (*Profile, error)
}

type apiV3 struct {
	client     *http.Client
	BaseURL    *url.URL
	UserAgent  string
	OAuthToken string
	IsCache    bool
	// API
	Categories *APICategories
	Entries    *APIEntries
	Markers    *APIMarkers
	*APIProfile
	Streams       *APIStreams
	Subscriptions *APISubscriptions
}

type Rate struct {
	Count int
	Limit int
	Reset time.Time
}

type Response struct {
	response *http.Response
	Rate
}

func GlobalResourceId(t GlobalResource, userId string) string {
	return fmt.Sprintf(t.String(), userId)
}

func ResourceId(t ResourceType, userId, identifier string) string {
	id := ""
	switch t {
	case RESOURCE_FEED:
		id = fmt.Sprintf("feed/%s", identifier)
	case RESOURCE_CATEGORY:
		id = fmt.Sprintf("user/%s/category/%s", userId, identifier)
	case RESOURCE_TAG:
		id = fmt.Sprintf("user/%s/tag/%s", userId, identifier)
	}
	return id
}

func NewAPI(httpClient *http.Client) API {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	baseURL, _ := url.Parse(baseURLCloud)
	baseURL.Path = version

	api := &apiV3{
		client:        httpClient,
		BaseURL:       baseURL,
		UserAgent:     "",
		OAuthToken:    "",
		IsCache:       false,
		Categories:    &APICategories{},
		Entries:       &APIEntries{},
		Markers:       &APIMarkers{},
		APIProfile:    &APIProfile{},
		Streams:       &APIStreams{},
		Subscriptions: &APISubscriptions{},
	}
	api.OAuthToken = os.Getenv("FEEDLY_ACCESS_TOKEN")

	api.IsCache = false
	api.Categories = &APICategories{api: api}
	api.Entries = &APIEntries{api: api}
	api.Markers = &APIMarkers{api: api}
	api.APIProfile = &APIProfile{api: api}
	api.Streams = &APIStreams{api: api}
	api.Subscriptions = &APISubscriptions{api: api}

	return api
}

func (c *apiV3) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	rel, err := url.Parse(version + "/" + urlStr)
	if err != nil {
		return nil, err
	}
	rawPath := rel.RawPath
	if rawPath == "" {
		rawPath = rel.Path
	}
	u := &url.URL{
		Host:     c.BaseURL.Host,
		Scheme:   c.BaseURL.Scheme,
		Opaque:   "//" + c.BaseURL.Host + "/" + rawPath,
		RawQuery: rel.RawQuery,
	}

	// u := c.BaseURL.ResolveReference(rel)
	// u.Path = version + "/" + urlStr

	// u := c.BaseURL
	// u.Path = rawPath
	// u.RawQuery = rel.RawQuery

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err = json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}
	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	if c.OAuthToken != "" {
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.OAuthToken))
	}

	if c.UserAgent != "" {
		req.Header.Add("User-Agent", c.UserAgent)
	}

	return req, nil
}

func newResponse(res *http.Response) *Response {
	r := &Response{response: res}

	if count := res.Header.Get(headerRateCount); count != "" {
		r.Rate.Count, _ = strconv.Atoi(count)
	}
	if limit := res.Header.Get(headerRateLimit); limit != "" {
		r.Rate.Limit, _ = strconv.Atoi(limit)
	}
	if reset := res.Header.Get(headerRateReset); reset != "" {
		if v, _ := strconv.ParseInt(reset, 10, 64); v != 0 {
			if t, err := time.Parse(http.TimeFormat, res.Header.Get("Date")); err == nil {
				r.Rate.Reset = t.Add(time.Duration(v * 1000000000))
			}
		}
	}

	return r
}

func (c *apiV3) Do(req *http.Request, v interface{}) (*Response, error) {
	rawPath := req.URL.RawPath
	if rawPath == "" {
		rawPath = req.URL.Path
	}
	dir := "./" + path.Dir(rawPath)
	base := path.Base(rawPath)
	q := req.URL.RawQuery
	if q != "" {
		q = "?" + q
	}
	p := path.Join(dir, base+url.QueryEscape(q)+".json")

	var res *http.Response
	var err error
	if c.IsCache /*&& req.Method == "GET"*/ {
		if _, err = os.Stat(dir); err != nil {
			if err = os.Mkdir(dir, 0755); err != nil {
				return nil, err
			}
		}
		var b []byte
		if b, err = ioutil.ReadFile(p); err == nil {
			res = &http.Response{
				Body: ioutil.NopCloser(bytes.NewBuffer(b)),
			}
		}
	}

	if res == nil {
		if res, err = c.client.Do(req); err != nil {
			return nil, err
		}
		if c.IsCache /* && req.Method == "GET" */ {
			var out *os.File
			if out, err = os.Create(p); err != nil {
				return nil, err
			}
			defer out.Close()
			bodyBytes, _ := ioutil.ReadAll(res.Body)
			if _, err = out.Write(bodyBytes); err != nil {
				return nil, err
			}
			res.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
		}
	}

	defer res.Body.Close()
	if res.StatusCode >= 400 {
		return nil, fmt.Errorf("bad response status code %d", res.StatusCode)
	}

	response := newResponse(res)

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			if _, err = io.Copy(w, res.Body); err != nil {
				return nil, err
			}
			// if err = json.NewEncoder(w).Encode(v); err != nil {
			// 	fmt.Println("aa %v", err)
			// 	return nil, err
			// }

		} else {
			err = json.NewDecoder(res.Body).Decode(v)
			if err == io.EOF {
				err = nil // ignore EOF errors caused by empty response body
			}
			// fmt.Printf("%v", v)
		}
	} else {
		// Debug
		var b []byte
		if b, err = ioutil.ReadAll(res.Body); err == nil {
			fmt.Println(string(b))
		}
	}

	return response, err
}

func addOptions(s string, opt interface{}) (string, error) {
	v := reflect.ValueOf(opt)
	if v.Kind() == reflect.Ptr && v.IsNil() {
		return s, nil
	}

	u, err := url.Parse(s)
	if err != nil {
		return s, err
	}

	qs, err := query.Values(opt)
	if err != nil {
		return s, err
	}

	u.RawQuery = qs.Encode()
	return u.String(), nil
}
