package lastfm

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
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
	url := fmt.Sprintf(lastfmAPIEndpoint+lastfmAPIAlbumInfoMethod, apiKey, url.QueryEscape(t.Artist), url.QueryEscape(t.Album))
	playlistmaker.LogInstance().Debug(fmt.Sprintf("GetAPIResult call url %s", url))
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var resultAPI map[string]interface{}
	errjson := json.Unmarshal(responseData, &resultAPI)
	if errjson != nil {
		playlistmaker.LogInstance().Warn("GetAPIResult error")
		playlistmaker.LogInstance().Warn("GetAPIResult responseData")
		playlistmaker.LogInstance().Warn(string(responseData))
		log.Fatal(errjson.Error())
	}
	playlistmaker.LogInstance().Debug(fmt.Sprintf("RESULT API CALL : %#v \n", resultAPI))
	playlistmaker.LogInstance().Debug("")
	playlistmaker.LogInstance().Debug("----------")
	playlistmaker.LogInstance().Debug("")
	albumObj, ok := resultAPI["album"]
	if !ok { //found album
		return nil
	}
	tracksObjs := albumObj.(map[string]interface{})["tracks"].(map[string]interface{})

	result = new(api.PlaylistAPIResult)
	result.Album = resultAPI["album"].(map[string]interface{})["name"].(string)
	trackMapsObj := tracksObjs["track"]
	trackMaps := trackMapsObj.([]interface{})
	if len(trackMaps) == 0 {
		return nil
	}
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
