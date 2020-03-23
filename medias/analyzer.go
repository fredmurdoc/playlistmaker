package medias

import (
	"log"
	"os"

	"github.com/dhowden/tag"
)

// TrackMetaData struct of metadata from track
type TrackMetaData struct {
	Filepath         string
	Name             string
	DurationInSecond uint64
	Order            int
	Album            string
	Artist           string
	Genre            string
}

//ExtractMetadata Test test function
func ExtractMetadata(f *os.File) (*TrackMetaData, error) {

	meta := TrackMetaData{}
	m, err := tag.ReadFrom(f)
	if err != nil {
		log.Fatal(err)
		return &meta, err
	}
	fileInfo, _ := f.Stat()
	meta.Filepath = fileInfo.Name()
	meta.Name = m.Title()
	meta.Album = m.Album()
	meta.Artist = m.Artist()
	meta.DurationInSecond = 0
	meta.Genre = m.Genre()
	trackNum, _ := m.Track()
	meta.Order = trackNum
	return &meta, err
}

//ToString get string representation
func (meta TrackMetaData) ToString() string {
	return "Duration : " + string(meta.DurationInSecond) +
		" - Title : " + meta.Name +
		" - Artist : " + meta.Artist +
		" - Album : " + meta.Album
}
