package api

import (
	"fmt"
	"github.com/ktr0731/go-fuzzyfinder"
	"log"
)

func FuzzyFinder() {
	idx, err := fuzzyfinder.FindMulti(
		tracks,
		func(i int) string {
			return tracks[i].FlagValue
		},
		fuzzyfinder.WithPreviewWindow(func(i, w, h int) string {
			if i == -1 {
				return ""
			}
			return fmt.Sprintf("Track: %s (%s)\nAlbum: %s",
				tracks[i].command,
				tracks[i].Flag,
				tracks[i].FlagValue)
		}))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("selected: %v\n", idx)
}
