package mock

import (
	"github.com/fredmurdoc/playlistmaker"
	"github.com/fredmurdoc/playlistmaker/api"
)

//Mock struct
type Mock struct {
	api.PlaylistAPIProviderInterface
}

//GetAPIResult result of mocks
func (m *Mock) GetAPIResult(t *playlistmaker.Track) (result *api.PlaylistAPIResult) {
	result = new(api.PlaylistAPIResult)

	result.Album = "AlbumTest"
	tracks := []api.TrackAPIResult{
		api.TrackAPIResult{
			Artist: "ArtisTest",
			Title:  "Title01",
			Order:  1,
			Length: 999,
		},
		api.TrackAPIResult{
			Artist: "ArtisTest",
			Title:  "Title02",
			Order:  2,
			Length: 999,
		},
		api.TrackAPIResult{
			Artist: "ArtisTest",
			Title:  "Title03",
			Order:  3,
			Length: 999,
		},
		api.TrackAPIResult{
			Artist: "ArtisTest",
			Title:  "Title04",
			Order:  4,
			Length: 999,
		},
	}
	result.Tracks = tracks
	return result

}
