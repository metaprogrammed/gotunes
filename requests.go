package gotunes

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type ItunesSearchRequest struct {
	Term      string `json:"term"`
	Country   string `json:"country"`
	Media     string `json:"media"`
	Entity    string `json:"entity"`
	Attribute string `json:"attribute"`
	Limit     string `json:"string"`
	Version   string `json:"version"`
	Explicit  bool   `json:"explicit"`
}

type ItunesFindRequest struct {
	ItunesId    string `json:"itunes_id"`
	AmgArtistId string `json:"amg_artist_id"`
	AmgAlbumId  string `json:"amg_album_id"`
	AmgVideoId  string `json:"amg_video_id"`
	Entity      string `json:"entity"`
	Isbn        string `json:"isbn"`
	Upc         string `json:"upc"`
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Search URL
func SearchUrl(request ItunesSearchRequest) string {
	Url, err := url.Parse("https://itunes.apple.com")
	if err != nil {
		log.Println("There was an error parsing the iTunes API Endpoint")
	}
	Url.Path += "/search"
	parameters := url.Values{}
	addParameter := func(key String, value String) {
		if value != "" {
			parameters.Add(key, value)
		}
	}
	addParameter("term", request.Term)
	addParameter("country", request.Country)
	addParameter("media", request.Media)
	addParameter("entity", request.Entity)
	addParameter("attribute", request.Attribute)
	addParameter("limit", request.Limit)
	if request.Explicit == false {
		parameters.Add("explicit", "no")
	}
	Url.RawQuery = parameters.Encode()
	return Url.String()
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Find URL
func FindUrl(request ItunesFindRequest) string {
	Url, err := url.Parse("http://itunes.apple.com")
	if err != nil {
		log.Println("There was an error parsing the iTunes API Endpoint")
	}
	Url.Path += "/lookup"
	parameters := url.Values{}
	addParameter := func(key String, value String) {
		if value != "" {
			parameters.Add(key, value)
		}
	}
	addParameter("id", request.ItunesId)
	addParameter("amgArtistId", request.AmgArtistId)
	addParameter("amgAlbumId", request.AmgAlbumId)
	addParameter("amgVideoId", request.AmgVideoId)
	addParameter("entity", request.Entity)
	addParameter("isbn", request.Isbn)
	addParameter("upc", request.Upc)
	Url.RawQuery = parameters.Encode()
	return Url.String()
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Make search request
func ItunesSearch(request ItunesSearchRequest) ItunesResponse {
	var url = SearchUrl(request)
	res, err := http.Get(url)
	if err != nil {
		log.Println("An error occurred when making the request: ", err)
	}
	defer res.Body.Close()
	contents, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println("An error occurred when getting the contents of the response: ", err)
	}
	var response ItunesResponse
	jsonerr := json.Unmarshal(contents, &response)
	if jsonerr != nil {
		log.Println("An error occurred unmarshaling the JSON: ", jsonerr)
	}
	return response
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Make find request
func ItunesFind(request ItunesFindRequest) ItunesResponse {
	var url = FindUrl(request)
	res, err := http.Get(url)
	if err != nil {
		log.Println("An error occurred when making the request: ", err)
	}
	defer res.Body.Close()
	contents, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println("An error occurred when getting the contents of the response: ", err)
	}
	var response ItunesResponse
	jsonerr := json.Unmarshal(contents, &response)
	if jsonerr != nil {
		log.Println("An error occurred unmarshaling the JSON: ", jsonerr)
	}
	return response
}
