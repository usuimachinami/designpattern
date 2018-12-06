package iterator

import (
	"github.com/usuimachinami/designpattern/iterator/entity"
)

// AlbumStatefulIterator contains information about an iterator's current state.
type AlbumStatefulIterator struct {
	current int
	album   entity.Album
}

// Next returns if next Song exists or not.
func (ai *AlbumStatefulIterator) Next() bool {
	ai.current++
	if ai.current >= len(ai.album.Songs) {
		return false
	}

	return true
}

// Song returns a current Song.
func (ai *AlbumStatefulIterator) Song() *entity.Song {
	return &ai.album.Songs[ai.current]
}

// NewAlbumStatefulIterator returns a stateful iterator for Album
func NewAlbumStatefulIterator(a entity.Album) *AlbumStatefulIterator {
	return &AlbumStatefulIterator{
		current: -1,
		album:   a,
	}
}
