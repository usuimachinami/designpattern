package iterator

import (
	"github.com/usuimachinami/designpattern/iterator/entity"
)

// AlbumCallbackIterator execute callback on each Songs.
func AlbumCallbackIterator(a *entity.Album, cb func(s *entity.Song)) {
	for _, s := range a.Songs {
		cb(&s)
	}
}
