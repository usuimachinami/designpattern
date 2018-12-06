package entity

import (
	"fmt"
	"time"
)

// Song contains information about an song.
type Song struct {
	Name     string
	Duration time.Duration
}

// Println prints the song name and duration in stdout with '\n'
func (s *Song) Println() {
	fmt.Printf("\t%s %f sec.\n", s.Name, s.Duration.Seconds())
}

// Album contains Songs.
type Album struct {
	Name  string
	Songs []Song
}
