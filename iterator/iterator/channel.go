package iterator

import (
	"github.com/usuimachinami/designpattern/iterator/entity"
)

// AlbumChannelIterator returns a channel for iteration.
func AlbumChannelIterator(a *entity.Album) <-chan *entity.Song {
	ch := make(chan *entity.Song)

	go func() {
		defer close(ch)
		for _, s := range a.Songs {
			song := s
			ch <- &song
		}
	}()

	return ch
}
