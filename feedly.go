package feedly

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
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
	GlobalMust GlobalResource = 1 << iota
	GlobalAll
	GlobalUncategorized
	GlobalRead
	GlobalSaved
)

func (t GlobalResource) String() string {
	switch t {
	case GlobalMust:
		return "user/%s/category/global.must"
	case GlobalAll:
		return "user/%s/category/global.all"
	case GlobalUncategorized:
		return "user/%s/category/global.uncategorized"
	case GlobalRead:
		return "user/%s/tag/global.read"
	case GlobalSaved:
		return "user/%s/tag/global.saved"
	}
	return ""
}

type ResourceType int

const (
	ResourceFeed ResourceType = 1 << iota
	ResourceCategory
	ResourceTag
)

type API interface {
	CollectionsCreate(context.Context, *CollectionCreate) (Collections, error)
	CollectionsFeedsDelete(context.Context, string, string) error
	CollectionsFeedsMDelete(context.Context, string, CollectionFeedDeletes) error
	CollectionsGet(context.Context, string) (Collections, error)
	CollectionsList(context.Context) (Collections, error)
	EntriesGet(context.Context, string) (Entries, error)
	FeedsGet(context.Context, string) (*Feed, error)
	FeedsMGet(context.Context, []string) (Feeds, error)
	MarkersCounts(context.Context) (*Marker, error)
	MarkersReads(context.Context, *MarkersReadsOptions) (*MarkersReads, error)
	ProfileGet(context.Context) (*Profile, error)
	StreamsContents(context.Context, string, *StreamOptions) (*StreamContents, error)
	StreamsIDs(context.Context, string, *StreamOptions) (*StreamIDs, error)
	SubscriptionsGet(context.Context) (Subscriptions, error)
	TagsList(context.Context) (Tags, error)
}

type apiV3 struct {
	client     *http.Client
	BaseURL    *url.URL
	UserAgent  string
	OAuthToken string
	IsCache    bool
	// API
	*APICollections
	*APIEntries
	*APIFeeds
	*APIMarkers
	*APIProfile
	*APIStreams
	*APISubscriptions
	*APITags
}

type Rate struct {
	Count int
	Limit int
	Reset time.Time
}

type Response struct {
	response *http.Response
	Rate     *Rate
}

func GlobalResourceID(t GlobalResource, userID string) string {
	return fmt.Sprintf(t.String(), userID)
}

func ResourceID(t ResourceType, userID, identifier string) (id string) {
	switch t {
	case ResourceFeed:
		id = fmt.Sprintf("feed/%s", identifier)
	case ResourceCategory:
		id = fmt.Sprintf("user/%s/category/%s", userID, identifier)
	case ResourceTag:
		id = fmt.Sprintf("user/%s/tag/%s", userID, identifier)
	}
	return
}

func NewAPI(httpClient *http.Client) API {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	baseURL, _ := url.Parse(baseURLCloud)
	baseURL.Path = version
	api := &apiV3{
		client:           httpClient,
		BaseURL:          baseURL,
		UserAgent:        "",
		OAuthToken:       "",
		IsCache:          false,
		APICollections:   &APICollections{},
		APIEntries:       &APIEntries{},
		APIFeeds:         &APIFeeds{},
		APIMarkers:       &APIMarkers{},
		APIProfile:       &APIProfile{},
		APIStreams:       &APIStreams{},
		APISubscriptions: &APISubscriptions{},
		APITags:          &APITags{},
	}
	api.OAuthToken = os.Getenv("FEEDLY_ACCESS_TOKEN")

	api.IsCache = false
	api.APICollections = &APICollections{api: api}
	api.APIEntries = &APIEntries{api: api}
	api.APIFeeds = &APIFeeds{api: api}
	api.APIMarkers = &APIMarkers{api: api}
	api.APIProfile = &APIProfile{api: api}
	api.APITags = &APITags{api: api}
	api.APIStreams = &APIStreams{api: api}
	api.APISubscriptions = &APISubscriptions{api: api}

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
		if err = json.NewEncoder(buf).Encode(body); err != nil {
			return nil, err
		}
	}
	req, err := http.NewRequest(method, u.String(), buf)
	// fmt.Printf("%s", buf)
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
	r := &Response{response: res, Rate: &Rate{}}
	if count := res.Header.Get(headerRateCount); count != "" {
		r.Rate.Count, _ = strconv.Atoi(count)
	}
	if limit := res.Header.Get(headerRateLimit); limit != "" {
		r.Rate.Limit, _ = strconv.Atoi(limit)
	}
	if reset := res.Header.Get(headerRateReset); reset != "" {
		const base = 10
		const bitSize = 64
		if v, _ := strconv.ParseInt(reset, base, bitSize); v != 0 {
			if t, err := time.Parse(http.TimeFormat, res.Header.Get("Date")); err == nil {
				const num = 1000000000
				r.Rate.Reset = t.Add(time.Duration(v * num))
			}
		}
	}

	return r
}

func (c *apiV3) Do(req *http.Request, v interface{}) (*Response, error) {
	var rawPath, dir, base, q string
	if rawPath = req.URL.RawPath; rawPath == "" {
		rawPath = req.URL.Path
	}
	dir = "./" + path.Dir(rawPath)
	base = path.Base(rawPath)
	if q = req.URL.RawQuery; q != "" {
		q = "?" + q
	}
	p := path.Join(dir, base+url.QueryEscape(q)+".json")
	var res *http.Response
	var err error
	if c.IsCache /*&& req.Method == "GET"*/ {
		if _, err = os.Stat(dir); err != nil {
			const perm = 0755
			if err = os.Mkdir(dir, perm); err != nil {
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
			defer res.Body.Close()
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
	if res.StatusCode >= http.StatusBadRequest {
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
			if errors.Is(err, io.EOF) {
				err = nil // ignore EOF errors caused by empty response body
			}
			// fmt.Printf("%v", v)
		}
	} else {
		// Debug
		var b []byte
		if b, err = ioutil.ReadAll(res.Body); err == nil {
			log.Println(string(b))
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
