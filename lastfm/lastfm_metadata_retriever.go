package lastfm

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var apiKey string = "ISMYSECRET"

const (
	lastfmAPIEndpoint          = "http://ws.audioscrobbler.com/2.0/?"
	lastfmAPITrackSearchMethod = "method=track.search&track=%s&api_key=%s&format=json"
)

// TrackMetaData struct of metadata from track
type TrackMetaData struct {
	filepath         string
	name             string
	durationInSecond uint64
	order            uint8
	album            string
	artist           string
	genre            string
}
type lastfmImage struct {
	Text string `json:"#text"`
	size string
}

type lastfmTrack struct {
	name       string
	artist     string
	url        string
	streamable string
	listeners  string
	image      []lastfmImage
}
type lastfmTrackSearchMatchItem struct {
	track []lastfmTrack
}

type openSearchQuery struct {
	Text      string `json:"#text"`
	role      string
	startPage string
}
type lastfmTrackSearchResultItem struct {
	OpensearchQuery        openSearchQuery `json:"opensearch:Query"`
	OpensearchTotalResults uint64          `json:"opensearch:totalResults"`
	OpensearchStartIndex   uint64          `json:"opensearch:startIndex"`
	OpensearchItemsPerPage uint64          `json:"opensearch:itemsPerPage"`
	trackmatches           lastfmTrackSearchMatchItem
}
type lastfmTrackSearchResponse struct {
	results lastfmTrackSearchResultItem
}

// GetMetadataFromMedia retrieve metadata from api
func GetMetadataFromMedia(track string) (metadata TrackMetaData) {
	url := fmt.Sprintf(lastfmAPIEndpoint+lastfmAPITrackSearchMethod, track, apiKey)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(string(responseData)[0:1000])
	//var jsonObj lastfmTrackSearchResponse
	var result map[string]interface{}

	errjson := json.Unmarshal(responseData, &result)
	if errjson != nil {
		log.Fatal(errjson.Error())
	}
	results := result["results"].(map[string]interface{})
	trackmatches := results["trackmatches"].(map[string]interface{})
	trackObj := trackmatches["track"].([]interface{})
	fmt.Printf("%#v", trackObj)

	//fmt.Println("opensearchItemsPerPage " + jsonObj.results.OpensearchQuery.Text)
	//fmt.Printf("%#v", jsonObj.results)
	return
}
