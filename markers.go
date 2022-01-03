package feedly

type APIMarkers struct {
	api *apiV3
}

type Marker struct {
	UnreadCounts []UnreadCount `json:"unreadcounts"`
	Updated      int64         `json:"updated"`
}

type UnreadCount struct {
	Count   int    `json:"count"`
	Id      string `json:"id"`
	Updated int64  `json:"updated"`
}

type MarkersReads struct {
	Entries []string           `json:"entries"`
	Feeds   []MarkersReadsFeed `json:"feeds"`
	Updated int64              `json:"updated"`
}

type MarkersReadsFeed struct {
	Id   string `json:"id"`
	AsOf int64  `json:"asOf"`
}

type MarkersReadsOptions struct {
	NewerThan int64 `url:"newerThan,omitempty"`
}

func (a *APIMarkers) Counts() (*Marker, *Response, error) {
	rel := "markers/counts"

	req, err := a.api.NewRequest("GET", rel, nil)
	if err != nil {
		return nil, nil, err
	}

	marker := new(Marker)

	res, err := a.api.Do(req, marker)
	if err != nil {
		return nil, res, err
	}

	return marker, res, nil
}

func (a *APIMarkers) Reads(opt *MarkersReadsOptions) (*MarkersReads, *Response, error) {
	rel := "markers/reads"
	rel, err := addOptions(rel, opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := a.api.NewRequest("GET", rel, nil)
	if err != nil {
		return nil, nil, err
	}

	markersReads := new(MarkersReads)

	res, err := a.api.Do(req, markersReads)
	if err != nil {
		return nil, res, err
	}

	return markersReads, res, nil
}
