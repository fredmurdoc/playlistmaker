package medias

import (
	"strconv"
	"testing"

	"github.com/fredmurdoc/playlistmaker/medias"
)

//TestAppendTrack
func TestAppendTrack(m *testing.T) {
	playlist1 := medias.Playlist{}
	track := medias.TrackMetaData{Filepath: "myFilepath", Name: "myTitle", Artist: "myArtist"}
	playlist1.AppendTrack(&track)
	track = medias.TrackMetaData{Filepath: "myFilepath2", Name: "myTitle2", Artist: "myArtist2"}
	playlist1.AppendTrack(&track)
	got := playlist1.Length()
	expected := 2
	if expected != got {
		m.Fatal("got :" + strconv.Itoa(got) + ", expected :" + strconv.Itoa(expected))
	}
}

//TestToString
func TestToString(m *testing.T) {
	playlist2 := medias.Playlist{}
	track := medias.TrackMetaData{Filepath: "myFilepath3", Name: "myTitle3", Artist: "myArtist3"}
	playlist2.AppendTrack(&track)
	track = medias.TrackMetaData{Filepath: "myFilepath4", Name: "myTitle4", Artist: "myArtist4"}
	playlist2.AppendTrack(&track)
	got := playlist2.ToString()
	// Output:
	expected := `#EXTM3U
#EXTINF:0, myArtist3 - myTitle3
myFilepath3
#EXTINF:0, myArtist4 - myTitle4
myFilepath4
`
	if expected != got {
		m.Fatal("got :" + got + ", expected :" + expected)
	}
}
