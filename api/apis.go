package api

import (
	"github.com/fredmurdoc/playlistmaker"
)

//PlaylistAPIProviderInterface interface for all api providers
type PlaylistAPIProviderInterface interface {
	GetAPIResult(t *playlistmaker.Track) *PlaylistAPIResult
}

//PlaylistAPIResult result of api interfacecall
type PlaylistAPIResult struct {
	Album  string
	Tracks []TrackAPIResult
}

//TrackAPIResult item of album
type TrackAPIResult struct {
	Artist string
	Title  string
	Order  int
	Length int
}

//GetAlbumPlaylistFromAPIProviderByNameAndArtist retrieve playlist from track
func GetAlbumPlaylistFromAPIProviderByNameAndArtist(t *playlistmaker.Track, api PlaylistAPIProviderInterface) *playlistmaker.Playlist {
	playlist := new(playlistmaker.Playlist)
	results := api.GetAPIResult(t)
	playlist = getPlaylistEntriesFromAPIResults(results)
	return playlist
}

//getPlaylistEntriesFromAPIResults:  parse les resultats de l'API
func getPlaylistEntriesFromAPIResults(results *PlaylistAPIResult) *playlistmaker.Playlist {

	playlist := new(playlistmaker.Playlist)
	for _, result := range results.Tracks {
		pe := new(playlistmaker.PlaylistEntry)
		t := new(playlistmaker.Track)
		t.Title = result.Title
		t.Album = results.Album
		t.Artist = result.Artist
		pe.Order = result.Order
		pe.Length = result.Length
		t.FileName = ""
		pe.Track = t
		playlist.Entries = append(playlist.Entries, pe)
	}
	return playlist
}
