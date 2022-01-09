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
	cid := fmt.Sprintf("user/%s/category/Feedly", profile.ID)
	var collections feedly.Collections
	if err = api.CollectionsFeedsMDelete(ctx, cid, feedly.CollectionFeedDeletes{
		{ID: "feed/http://feeds.feedburner.com/design-milk"},
	}); err != nil {
		panic(err)
	}
	if collections, err = api.CollectionsGet(ctx, cid); err != nil {
		panic(err)
	}
	fmt.Printf("%s", collections)
}
