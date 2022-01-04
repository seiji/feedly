package main

import (
	"context"
	"fmt"

	"github.com/seiji/feedly"
)

func main() {
	api := feedly.NewAPI(nil)
	ctx := context.Background()
	ids := []string{
		"feed/http://feeds.engadget.com/weblogsinc/engadget",
		"feed/http://www.yatzer.com/feed/index.php",
	}
	got, err := api.FeedsMGet(ctx, ids)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	fmt.Printf("%s", got)
}
