package main

import (
	"fmt"
	"time"

	"github.com/usuimachinami/designpattern/iterator/entity"
	"github.com/usuimachinami/designpattern/iterator/iterator"
)

func main() {
	a := entity.Album{
		Name: "New Album",
		Songs: []entity.Song{
			{
				Name:     "Song 1",
				Duration: time.Duration(180 * time.Second),
			},
			{
				Name:     "Song 2",
				Duration: time.Duration(300 * time.Second),
			},
			{
				Name:     "Song 3",
				Duration: time.Duration(90 * time.Second),
			},
		},
	}

	fmt.Println("Iterator Pattern\n")

	execStatefulPattern(a)

	execCallbackPattern(a)

	execClosurePattern(a)

	execChannelPattern(a)
}

func execStatefulPattern(a entity.Album) {
	fmt.Println("Stateful")
	fmt.Println(a.Name)
	asi := iterator.NewAlbumStatefulIterator(a)
	for asi.Next() {
		asi.Song().Println()
	}
	fmt.Println()
}

func execCallbackPattern(a entity.Album) {
	fmt.Println("Callback")
	fmt.Println(a.Name)
	cb := func(s *entity.Song) {
		s.Println()
	}
	iterator.AlbumCallbackIterator(&a, cb)
	fmt.Println()
}

func execClosurePattern(a entity.Album) {
	fmt.Println("Closure")
	fmt.Println(a.Name)
	aci := iterator.AlbumClosureIterator(&a)
	for s, ok := aci(); ok; s, ok = aci() {
		s.Println()
	}
	fmt.Println()
}

func execChannelPattern(a entity.Album) {
	fmt.Println("Channel")
	fmt.Println(a.Name)
	for s := range iterator.AlbumChannelIterator(&a) {
		s.Println()
	}
	fmt.Println()
}
