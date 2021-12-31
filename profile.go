package feedly

import (
	"encoding/json"
	"net/http"
)

type APIProfile struct {
	client *Client
}

type Login struct {
	FullName   string `json:"fullName"`
	Id         string `json:"id"`
	Picture    string `json:"picture"`
	Provider   string `json:"provider"`
	ProviderId string `json:"providerId"`
	Verified   bool   `json:"verified"`
}
type PaymentProviderId struct {
	Paypal string `json:"Paypal"`
}

type PaymentSubscriptionId struct {
	Paypal string `json:"Paypal"`
}

type Profile struct {
	Client                      string                 `json:"client"`
	DropboxConnected            bool                   `json:"dropboxConnected"`
	Email                       string                 `json:"email"`
	EvernoteConnected           bool                   `json:"evernoteConnected"`
	FacebookConnected           bool                   `json:"facebookConnected"`
	FamilyName                  string                 `json:"familyName"`
	FullName                    string                 `json:"fullName"`
	Gender                      string                 `json:"gender"`
	GivenName                   string                 `json:"givenName"`
	Google                      string                 `json:"google"`
	Id                          string                 `json:"id"`
	Locale                      string                 `json:"locale"`
	Logins                      []Login                `json:"logins"`
	PaymentProviderId           *PaymentProviderId     `json:"paymentProviderId"`
	PaymentSubscriptionId       *PaymentSubscriptionId `json:"paymentSubscriptionId"`
	Picture                     string                 `json:"picture"`
	PocketConnected             bool                   `json:"pocketConnected"`
	Product                     string                 `json:"product"`
	ProductExpiration           int64                  `json:"productExpiration"`
	SubscriptionPaymentProvider string                 `json:"subscriptionPaymentProvider"`
	SubscriptionStatus          string                 `json:"subscriptionStatus"`
	TwitterConnected            bool                   `json:"twitterConnected"`
	UpgradeDate                 int64                  `json:"upgradeDate"`
	Wave                        string                 `json:"wave"`
	WindowsLiveConnected        bool                   `json:"windowsLiveConnected"`
	WordPressConnected          bool                   `json:"wordPressConnected"`
}

func (a Profile) String() string {
	e, err := json.Marshal(a)
	if err != nil {
		panic(err)
	}
	return string(e)
}

func (a *APIProfile) Get() (p *Profile, err error) {
	var req *http.Request
	if req, err = a.client.NewRequest("GET", "profile", nil); err != nil {
		return nil, err
	}
	p = new(Profile)
	// var res *Response
	if _, err = a.client.Do(req, p); err != nil {
		return nil, err
	}
	return p, nil
}

func (a *APIProfile) Update() (*Profile, *Response, error) {
	return nil, nil, nil
}
