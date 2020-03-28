package lastfm

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/fredmurdoc/playlistmaker"
	"github.com/fredmurdoc/playlistmaker/api"
)

var apiKey string = "4ce571ae066f563e18a5da1048545658"

const (
	lastfmAPIEndpoint          = "http://ws.audioscrobbler.com/2.0/?"
	lastfmAPITrackSearchMethod = "method=track.search&track=%s&api_key=%s&format=json"
	lastfmAPIAlbumSearchMethod = "method=album.search&album=%s&api_key=%s&format=json"
	lastfmAPIAlbumInfoMethod   = "method=album.getinfo&api_key=%s&artist=%s&album=%s&format=json"
)

// TrackMetaData struct of metadata from track
type TrackMetaData struct {
	Filepath         string
	Name             string
	DurationInSecond uint64
	Order            uint8
	Album            string
	Artist           string
	Genre            string
}

// GetAPIResult retrieve result from api
func GetAPIResult(t *playlistmaker.Track) (result *api.PlaylistAPIResult) {
	url := fmt.Sprintf(lastfmAPIEndpoint+lastfmAPIAlbumInfoMethod, apiKey, t.Artist, t.Album)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	result = new(api.PlaylistAPIResult)
	//fmt.Println(string(responseData)[0:1000])
	//var jsonObj lastfmTrackSearchResponse
	var resultAPI map[string]interface{}
	//fmt.Printf("%#v \n", string(responseData))
	errjson := json.Unmarshal(responseData, &resultAPI)
	if errjson != nil {
		log.Fatal(errjson.Error())
	}
	fmt.Printf("%#v \n", resultAPI)
	fmt.Println()
	fmt.Println("----------")
	fmt.Println()

	tracksObjs := resultAPI["album"].(map[string]interface{})["tracks"].(map[string]interface{})

	trackMapsObj := tracksObjs["track"]
	trackMaps := trackMapsObj.([]interface{})
	for indexTrack, trackMap := range trackMaps {
		fmt.Printf("TRACK INDEX %v \n", indexTrack)
		fmt.Printf("%#v \n", trackMap)
		fmt.Println("--------------")
	}

	return result
}
