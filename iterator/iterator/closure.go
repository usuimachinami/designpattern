package iterator

import (
	"github.com/usuimachinami/designpattern/iterator/entity"
)

// AlbumClosureIterator returns a generator function as a iterator.
func AlbumClosureIterator(a *entity.Album) func() (*entity.Song, bool) {
	current := 0
	totalLen := len(a.Songs)

	return func() (*entity.Song, bool) {
		defer func() {
			current++
		}()

		if current < totalLen {
			return &a.Songs[current], true
		} else {
			return nil, false
		}
	}
}
