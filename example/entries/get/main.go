package main

import (
	"context"
	"fmt"

	"github.com/seiji/feedly"
)

func main() {
	api := feedly.NewAPI(nil)
	ctx := context.Background()
	got, err := api.EntriesGet(ctx, "Xne8uW/IUiZhV1EuO2ZMzIrc2Ak6NlhGjboZ+Yk0rJ8=_1523699cbb3:2aa0463:e47a7aef")
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	fmt.Printf("%s", got)
}
