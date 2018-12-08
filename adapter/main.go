package main

import (
	"fmt"
	"time"

	"github.com/usuimachinami/designpattern/adapter/adapter"
	"github.com/usuimachinami/designpattern/adapter/client"
	"github.com/usuimachinami/designpattern/adapter/entity"
)

func main() {
	s := entity.Song{
		Name:     "Song 1",
		Duration: 100 * time.Second,
	}
	fmt.Print("Adapter Pattern\n\n")

	fmt.Println("Embed")
	embed := adapter.SongEmbedAdapter{
		Song: s,
	}
	fmt.Println(client.SongInfoClient(embed))
	fmt.Println()

	fmt.Println("Transfer")
	transfer := adapter.SongEmbedAdapter{
		Song: s,
	}
	fmt.Println(client.SongInfoClient(transfer))
	fmt.Println()
}
