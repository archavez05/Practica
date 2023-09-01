package main

type Song struct {
	id       int     `json:"id"`
	originId string  `json:originId`
	name     string  `json:"name"`
	artist   string  `json:"artist"`
	duration string  `json:"duration"`
	album    string  `json:"album"`
	artwork  string  `json:"artwork"`
	price    float64 `json:"price"`
	origin   string  `json:"origin"`
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
	resultCount int                `json:"resultCount"`
	results     []SongRestResponse `json:"results"`
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
