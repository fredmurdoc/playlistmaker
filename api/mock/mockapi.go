package mock

import (
	"github.com/fredmurdoc/playlistmaker"
	"github.com/fredmurdoc/playlistmaker/api"
)

//Mock interface
type Mock interface {
	PlaylistAPIProviderInterface
}

//GetAPIResult result of mocks
func (m *Mock) GetAPIResult(t *playlistmaker.Track) (result *api.PlaylistAPIResult) {
	result = new(api.PlaylistAPIResult)

	result.Album = "AlbumTest"
	tracks := []api.TrackAPIResult{
		api.TrackAPIResult{
			Artist: "ArtisTest",
			Order:  1,
			Length: 999,
		},
		api.TrackAPIResult{
			Artist: "ArtisTest",
			Order:  2,
			Length: 999,
		},
		api.TrackAPIResult{
			Artist: "ArtisTest",
			Order:  3,
			Length: 999,
		},
		api.TrackAPIResult{
			Artist: "ArtisTest",
			Order:  4,
			Length: 999,
		},
	}
	result.Tracks = tracks
	return result

}
