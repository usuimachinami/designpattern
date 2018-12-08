package adapter

import (
	"fmt"

	"github.com/usuimachinami/designpattern/adapter/entity"
)

// SongEmbedAdapter has a embed song.
type SongEmbedAdapter struct {
	entity.Song
}

// Show returns information about the embed song
func (sa SongEmbedAdapter) Show() string {
	return fmt.Sprintf("%s / %d", sa.GetName(), sa.GetDuration())
}
