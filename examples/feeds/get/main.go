package main

import (
	"context"
	"fmt"

	"github.com/seiji/feedly"
)

func main() {
	api := feedly.NewAPI(nil)
	ctx := context.Background()
	id := "feed/http://feeds.engadget.com/weblogsinc/engadget"
	got, err := api.FeedsGet(ctx, id)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	fmt.Printf("%s", got)
}
