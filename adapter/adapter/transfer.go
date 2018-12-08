package adapter

import (
	"fmt"

	"github.com/usuimachinami/designpattern/adapter/entity"
)

// SongTransferAdapter has a reference of a song.
type SongTransferAdapter struct {
	s *entity.Song
}

// Show returns information about the embed song
func (sa SongTransferAdapter) Show() string {
	return fmt.Sprintf("%s / %d", sa.s.GetName(), sa.s.GetDuration())
}
