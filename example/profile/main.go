package main

import (
	"fmt"

	"github.com/seiji/feedly"
)

func main() {
	api := feedly.NewAPI(nil)
	profile, err := api.ProfileGet(nil)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	fmt.Printf("%s", profile)
}
