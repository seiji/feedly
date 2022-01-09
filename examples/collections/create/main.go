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
	var collections feedly.Collections
	if collections, err = api.CollectionsCreate(ctx, &feedly.CollectionCreate{
		Description: "from ci",
		Feeds: []feedly.CollectionFeedCreate{
			{
				ID: "feed/http://feeds.feedburner.com/design-milk",
			},
		},
		ID:    cid,
		Label: label,
	}); err != nil {
		panic(err)
	}
	fmt.Printf("%s", collections)
}
