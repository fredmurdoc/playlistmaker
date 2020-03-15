package medias

import (
	"testing"

	"github.com/fredmurdoc/playlistmaker/medias"
)

//TestAppendTrack
func TestAppendTrack(m *testing.T) {
	track := medias.TrackMetaData{filepath: "filepath", name: "Title", artist: "Artist"}
	medias.AppendTrack(track)
}
