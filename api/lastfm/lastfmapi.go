package lastfm

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

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

//LastFM struct
type LastFM struct {
	api.PlaylistAPIProviderInterface
}

// GetAPIResult retrieve result from api
func (m *LastFM) GetAPIResult(t *playlistmaker.Track) (result *api.PlaylistAPIResult) {
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
	playlistmaker.LogInstance().Debug(fmt.Sprintf("%#v \n", resultAPI))
	playlistmaker.LogInstance().Debug("")
	playlistmaker.LogInstance().Debug("----------")
	playlistmaker.LogInstance().Debug("")

	tracksObjs := resultAPI["album"].(map[string]interface{})["tracks"].(map[string]interface{})
	result.Album = resultAPI["album"].(map[string]interface{})["name"].(string)
	trackMapsObj := tracksObjs["track"]
	trackMaps := trackMapsObj.([]interface{})
	for _, trackMap := range trackMaps {

		t := api.TrackAPIResult{}
		t.Artist = trackMap.(map[string]interface{})["artist"].(map[string]interface{})["name"].(string)

		t.Title = trackMap.(map[string]interface{})["name"].(string)
		duration, errConv := strconv.Atoi(trackMap.(map[string]interface{})["duration"].(string))
		if errConv != nil {
			duration = 0
		}
		t.Length = duration

		order, errConv2 := strconv.Atoi(trackMap.(map[string]interface{})["@attr"].(map[string]interface{})["rank"].(string))
		if errConv2 != nil {
			order = 0
		}
		t.Order = order
		result.Tracks = append(result.Tracks, t)
	}

	return result
}
