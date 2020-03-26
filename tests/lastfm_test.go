package tests

import (
	"fmt"
	"testing"

	"github.com/fredmurdoc/playlistmaker/lastfm"
)

//TestAppendTrack
func TestGetAlbumMetadataFromMedia(m *testing.T) {
	fmt.Println(lastfm.GetAlbumMetadataFromName("Functional Arrhythmias"))
}
