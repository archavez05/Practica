package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.Handle("/getSOAPSongsByArtist", validateToken(http.HandlerFunc(getSOAPSongsByArtist)))
	myRouter.Handle("/getRestSongByArtistAndSong", validateToken(http.HandlerFunc(getRestSongByArtistAndSong)))
	myRouter.Handle("/createUser", validateToken(http.HandlerFunc(createUser))).Methods("POST")
	myRouter.HandleFunc("/login", login).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", myRouter))
}
