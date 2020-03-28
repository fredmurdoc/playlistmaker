package tests

import (
	"testing"

	"github.com/fredmurdoc/playlistmaker"

	"github.com/fredmurdoc/playlistmaker/api/mock"
)

func TestGetAlbumPlaylistFromNameAndArtist(m *testing.T) {
	mock := new(mock.Mock)
	track := new(playlistmaker.Track)
	Playlist playlist = mock.GetAlbumPlaylistFromNameAndArtist(track)
}
