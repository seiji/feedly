package main

import (
	"fmt"

	"github.com/seiji/feedly"
)

func main() {
	api := feedly.New(nil)
	// var profile *feedly.Profile
	// var res *feedly.Response
	// var err error
	profile, err := api.Profile.Get()
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	fmt.Printf("%s", profile)
}
