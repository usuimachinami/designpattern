package entity

import (
	"time"
)

// Song contains information about an song.
type Song struct {
	Name     string
	Duration time.Duration
}

// GetName returns a song name.
func (s *Song) GetName() string {
	return s.Name
}

// GetDuration returns song duration seconds.
func (s *Song) GetDuration() int {
	return int(s.Duration.Seconds())
}
