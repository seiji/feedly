package main

import (
	"context"
	"fmt"

	"github.com/seiji/feedly"
)

func main() {
	api := feedly.NewAPI(nil)
	ctx := context.Background()
	var err error
	var profile *feedly.Profile
	if profile, err = api.ProfileGet(ctx); err != nil {
		panic(err)
	}
	label := "Feedly"
	cid := fmt.Sprintf("user/%s/category/%s", profile.ID, label)
	var feeds feedly.Feeds
	if feeds, err = api.CollectionsFeedsMPut(ctx, cid, feedly.CollectionFeedCreates{
		{
			ID: "feed/http://feeds.feedburner.com/design-milk",
		},
	}); err != nil {
		panic(err)
	}
	fmt.Printf("%s", feeds)
}
