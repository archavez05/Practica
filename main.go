package main

import (
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	/*newSong := Song{name: "test",
		artist:   "test",
		duration: "test",
		album:    "test",
		artwork:  "test",
		price:    0.90,
		origin:   "test",
	}
	insertSong(newSong)
	getSongs()
	getSongByName("test")
	getSongByArtist("test")
	getSongByAlbum("test")
	getSongByArtist()*/
	handleRequests()
}
