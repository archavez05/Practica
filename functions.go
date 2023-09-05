package main

import (
	"database/sql"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func getRestSongByArtistAndSong(w http.ResponseWriter, r *http.Request) {

	resp, err := http.Get("https://itunes.apple.com/search?term=Eminem&limit=5")
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	var stream io.Reader = resp.Body

	fmt.Println(stream)
	var response RestResponse
	err = json.NewDecoder(stream).Decode(&response)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(response)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response.Results)
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
	responseSOAP := SOAPResponse{}
	xml.Unmarshal([]byte(bodyString), &responseSOAP)
	var songs []Song
	for _, value := range responseSOAP.SearchLyricResult {
		if value.Song == "" {
			break
		}
		song := Song{
			OriginId: string(value.TrackId),
			Name:     value.Song,
			Artist:   value.Artist,
			Origin:   "chartlyrics",
		}
		fmt.Println("Parseo -- ", song)
		insertSong(song)
		fmt.Println("Guardado -- ", song)
		songs = append(songs, song)
	}
	fmt.Println(len(songs))
	var response = GeneralResponse{
		Message: "Listado de Canciones",
		Songs:   songs,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

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

	resp, err := insert.Query(song.Name, song.Artist, song.Duration, song.Album, song.Artwork, song.Price, song.Origin)
	if err != nil {
		fmt.Println(err)
	}
	insert.Close()

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp)
	var newSong Song
	for resp.Next() {
		err = resp.Scan(&newSong.Id, &newSong.OriginId, &newSong.Name, &newSong.Artist, &newSong.Duration, &newSong.Album, &newSong.Price, &newSong.Origin)
		if err != nil {
			fmt.Println("ERROR en Guardado", err)
		}
	}
	fmt.Println("NUEVA CANCIÓN - - ", newSong)
}

/*
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
}*/
