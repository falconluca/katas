package dao

import (
	"log"
)

func TestDb() {
	InitMySQL()

	//SaveAlbum("Allen", "NBA", 80.25)
	//UpdateAlbum("Go in Action", 5)
	//DeleteAlbum(5)
	albums, _ := ListAllAlbums()
	for _, album := range albums {
		log.Printf("id: %v, title: %v", album.ID, album.Title)
	}
}
