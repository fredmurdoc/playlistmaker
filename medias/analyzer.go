package medias

import (
	"log"
	"os"

	"github.com/dhowden/tag"
)

// TrackMetaData struct of metadata from track
type TrackMetaData struct {
	filepath         string
	name             string
	durationInSecond uint64
	order            int
	album            string
	artist           string
	genre            string
}

//Test test function
func Test(f *os.File) (*TrackMetaData, error) {

	meta := TrackMetaData{}
	m, err := tag.ReadFrom(f)
	if err != nil {
		log.Fatal(err)
		return &meta, err
	}
	fileInfo, _ := f.Stat()
	meta.filepath = fileInfo.Name()
	meta.name = m.Title()
	meta.album = m.Album()
	meta.artist = m.Artist()
	meta.durationInSecond = 0
	meta.genre = m.Genre()
	trackNum, _ := m.Track()
	meta.order = trackNum
	return &meta, err
}

//ToString get string representation
func (meta TrackMetaData) ToString() string {
	return string(meta.durationInSecond) + " - " + meta.name + " - " + meta.artist
}
