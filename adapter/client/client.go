package client

import (
	"github.com/usuimachinami/designpattern/adapter/adapter"
)

// SongInfoClient is a client of songs.
func SongInfoClient(s adapter.Shower) string {
	return s.Show()
}
