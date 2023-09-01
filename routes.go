package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", testRest)
	myRouter.HandleFunc("/soap", getSOAPSongsByArtist)
	myRouter.HandleFunc("/rest", getRestSongByArtist)
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}
