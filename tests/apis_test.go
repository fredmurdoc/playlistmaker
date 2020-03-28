package tests

import (
	"fmt"
	"testing"

	"github.com/fredmurdoc/playlistmaker"

	"github.com/fredmurdoc/playlistmaker/api"
	"github.com/fredmurdoc/playlistmaker/api/lastfm"
	"github.com/fredmurdoc/playlistmaker/api/mock"
)

func TestMockGetAlbumPlaylistFromAPIProviderByNameAndArtist(m *testing.T) {
	mock := new(mock.Mock)
	track := new(playlistmaker.Track)
	playlist := api.GetAlbumPlaylistFromAPIProviderByNameAndArtist(track, mock)
	//fmt.Printf("%s", playlist.String())
	expected := 4
	got := len(playlist.Entries)
	if expected != got {
		m.Fatalf("expected %d got %d", expected, got)
	}
}

func TestLastFmGetAlbumPlaylistFromAPIProviderByNameAndArtist(m *testing.T) {
	mock := new(lastfm.LastFM)
	track := new(playlistmaker.Track)
	track.Album = "Bad"
	track.Artist = "Michael Jackson"
	playlist := api.GetAlbumPlaylistFromAPIProviderByNameAndArtist(track, mock)
	fmt.Printf("%s", playlist.String())
	expected := 10
	got := len(playlist.Entries)
	if expected != got {
		m.Fatalf("expected %d got %d", expected, got)
	}
}
