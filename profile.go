package feedly

import (
	"context"
	"encoding/json"
	"net/http"
)

type apiProfile struct {
	api *apiV3
}

type Login struct {
	FullName   string `json:"fullName"`
	ID         string `json:"id"`
	Picture    string `json:"picture"`
	Provider   string `json:"provider"`
	ProviderID string `json:"providerId"`
	Verified   bool   `json:"verified"`
}
type PaymentProviderID struct {
	Paypal string `json:"Paypal"`
}

type PaymentSubscriptionID struct {
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
	ID                          string                 `json:"id"`
	Locale                      string                 `json:"locale"`
	Logins                      []Login                `json:"logins"`
	PaymentProviderID           *PaymentProviderID     `json:"paymentProviderId"`
	PaymentSubscriptionID       *PaymentSubscriptionID `json:"paymentSubscriptionId"`
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

func (a *apiProfile) ProfileGet(ctx context.Context) (p *Profile, err error) {
	var req *http.Request
	if req, err = a.api.NewRequest(ctx, "GET", "profile", nil); err != nil {
		return nil, err
	}
	p = new(Profile)
	if _, err = a.api.Do(req, p); err != nil {
		return nil, err
	}
	return p, nil
}

// func (a *apiProfile) Update(ctx context.Context) (p *Profile, err error) {
// 	return
// }
