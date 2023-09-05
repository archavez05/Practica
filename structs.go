package main

import "github.com/golang-jwt/jwt/v5"

type Song struct {
	Id       int     `json:"id"`
	OriginId string  `json:"originId"`
	Name     string  `json:"name"`
	Artist   string  `json:"artist"`
	Duration string  `json:"duration"`
	Album    string  `json:"album"`
	Artwork  string  `json:"artwork"`
	Price    float64 `json:"price"`
	Origin   string  `json:"origin"`
}

type SongRestResponse struct {
	WrapperType            string  `json:"wrapperType"`
	Kind                   string  `json:"kind"`
	ArtistId               int     `json:"artistId"`
	CollectionId           int     `json:"collectionId"`
	TrackId                int     `json:"trackId"`
	ArtistName             string  `json:"artistName"`
	CollectionName         string  `json:"collectionName"`
	TrackName              string  `json:"trackName"`
	CollectionCensoredName string  `json:"collectionCensoredName"`
	TrackCensoredName      string  `json:"trackCensoredName"`
	ArtistViewUrl          string  `json:"artistViewUrl"`
	CollectionViewUrl      string  `json:"collectionViewUrl"`
	TrackViewUrl           string  `json:"trackViewUrl"`
	PreviewUrl             string  `json:"previewUrl"`
	ArtworkUrl30           string  `json:"artworkUrl30"`
	ArtworkUrl60           string  `json:"artworkUrl60"`
	ArtworkUrl100          string  `json:"artworkUrl100"`
	CollectionPrice        float64 `json:"collectionPrice"`
	TrackPrice             float64 `json:"trackPrice"`
	ReleaseDate            string  `json:"releaseDate"`
	CollectionExplicitness string  `json:"collectionExplicitness"`
	TrackExplicitness      string  `json:"trackExplicitness"`
	DiscCount              int     `json:"discCount"`
	DiscNumber             int     `json:"discNumber"`
	TrackCount             int     `json:"trackCount"`
	TrackNumber            int     `json:"trackNumber"`
	TrackTimeMillis        int     `json:"trackTimeMillis"`
	Country                string  `json:"country"`
	Currency               string  `json:"currency"`
	PrimaryGenreName       string  `json:"primaryGenreName"`
	IsStreamable           bool    `json:"isStreamable"`
}

type RestResponse struct {
	ResultCount int                `json:"resultCount"`
	Results     []SongRestResponse `json:"results"`
}

type SongSOAPResponse struct {
	TrackChecksum string `xml:"TrackChecksum"`
	TrackId       int    `xml:"TrackId"`
	LyricId       int    `xml:"LyricId"`
	SongUrl       string `xml:"SongUrl"`
	ArtistUrl     string `xml:"ArtistUrl"`
	Artist        string `xml:"Artist"`
	Song          string `xml:"Song"`
	SongRank      int    `xml:"SongRank"`
}

type SOAPResponse struct {
	SearchLyricResult []SongSOAPResponse `xml:SearchLyricResult`
}

type LoginRequest struct {
	User string
	Pass string
}

type User struct {
	Id    int
	User  string
	Pass  string
	Email string
}

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

type LoginResponse struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}

type GeneralResponse struct {
	Message string `json:"message"`
	Songs   []Song `json:"songs"`
}
