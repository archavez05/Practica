package main

import (
	"database/sql"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func getRestSongByArtist(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("https://itunes.apple.com/search?term=Eminem&limit=1")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	bodyString := string(bodyBytes)
	fmt.Println(bodyString)
	var response RestResponse
	err = json.Unmarshal(bodyBytes, &response)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%v \n", response)

}

func getSOAPSongsByArtist(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("http://api.chartlyrics.com/apiv1.asmx/SearchLyric?artist=Eminem&song=encore")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	bodyString := string(bodyBytes)
	fmt.Println(bodyString)
	response := SOAPResponse{}
	xml.Unmarshal([]byte(bodyString), &response)
	fmt.Printf("%v \n", response)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func testRest(w http.ResponseWriter, r *http.Request) {
	fmt.Println("API funcionando")
}

func insertSong(song Song) {
	db, err := sql.Open("mysql", "root:48821181Ap!@tcp(127.0.0.1:3306)/Practice")

	if err != nil {
		panic(err)
	}
	query := "INSERT INTO `songs` (name, artist, duration, album, artwork, price, origin) values (?,?,?,?,?,?,?)"
	insert, err := db.Prepare(query)
	if err != nil {
		fmt.Println(err)
	}

	resp, err := insert.Exec(song.name, song.artist, song.duration, song.album, song.artwork, song.price, song.origin)
	insert.Close()

	if err != nil {
		fmt.Println(err)
	}

	println(resp)
}

func findSongs() {
	var song Song
	db, err := sql.Open("mysql", "root:48821181Ap!@tcp(127.0.0.1:3306)/Practice")

	if err != nil {
		panic(err)
	}

	defer db.Close()

	results, err := db.Query("SELECT * FROM songs")
	if err != nil {
		panic(err.Error())
	}

	for results.Next() {

		err := results.Scan(
			&song.id,
			&song.name,
			&song.artist,
			&song.duration,
			&song.album,
			&song.artwork,
			&song.price,
			&song.origin)
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("%v \n", song)
	}
}

func findSongByName(name string) {
	var song Song
	db, err := sql.Open("mysql", "root:48821181Ap!@tcp(127.0.0.1:3306)/Practice")

	if err != nil {
		panic(err)
	}

	defer db.Close()

	results, err := db.Query("SELECT * FROM songs where name = ?", name)
	if err != nil {
		panic(err.Error())
	}

	for results.Next() {

		err := results.Scan(
			&song.id,
			&song.name,
			&song.artist,
			&song.duration,
			&song.album,
			&song.artwork,
			&song.price,
			&song.origin)
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("%v \n", song)
	}
}

func findSongByArtist(artist string) {
	var song Song
	db, err := sql.Open("mysql", "root:48821181Ap!@tcp(127.0.0.1:3306)/Practice")

	if err != nil {
		panic(err)
	}

	defer db.Close()

	results, err := db.Query("SELECT * FROM songs where artist = ?", artist)
	if err != nil {
		panic(err.Error())
	}

	for results.Next() {

		err := results.Scan(
			&song.id,
			&song.name,
			&song.artist,
			&song.duration,
			&song.album,
			&song.artwork,
			&song.price,
			&song.origin)
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("%v \n", song)
	}
}

func findSongByAlbum(album string) {
	var song Song
	db, err := sql.Open("mysql", "root:48821181Ap!@tcp(127.0.0.1:3306)/Practice")

	if err != nil {
		panic(err)
	}

	defer db.Close()

	results, err := db.Query("SELECT * FROM songs where album = ?", album)
	if err != nil {
		panic(err.Error())
	}

	for results.Next() {

		err := results.Scan(
			&song.id,
			&song.name,
			&song.artist,
			&song.duration,
			&song.album,
			&song.artwork,
			&song.price,
			&song.origin)
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("%v \n", song)
	}
}
