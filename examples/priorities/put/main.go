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
	var priorities feedly.Priorities
	if priorities, err = api.PrioritiesPut(ctx, feedly.PriorityCreate{
		Active:      false,
		ActiveUntil: 0,
		Label:       "test",
		Layers: []feedly.Filter{
			{
				Boards:        []feedly.Board{},
				ComplexFilter: false,
				Parts: []feedly.FilterPart{
					{
						ID:    "nlp/f/entity/wd:Q30",
						Label: "Machine Learning",
						Text:  "",
					},
				},
				Salience:   "about",
				Severities: []string{},
				Type:       "matches",
			},
		},
		StreamIDs: []string{feedly.GlobalCategoryAll.ID(profile.ID)},
	}); err != nil {
		panic(err)
	}
	fmt.Printf("%s", priorities)
}
