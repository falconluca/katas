package dao

import (
	"fmt"
	"log"
)

func TestDb() {
	InitMySQL()

	albums, err := AlbumsByArtist("John Coltrane")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Albums found: %v\n", albums)
}
