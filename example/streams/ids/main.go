package main

import (
	"context"
	"fmt"

	"github.com/seiji/feedly"
)

func main() {
	api := feedly.NewAPI(nil)
	ctx := context.Background()
	id := "feed/https://www.theverge.com/rss/full.xml"
	got, err := api.StreamsIDs(ctx, id, nil)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	fmt.Printf("%s", got)
}
