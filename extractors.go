package playlistmaker

import (
	"log"
	"os"

	"github.com/dhowden/tag"
)

//ExtractMetadataToTrack : extract metadatafrom media file
func ExtractMetadataToTrack(f *os.File, t *Track) error {
	m, err := tag.ReadFrom(f)
	if err != nil {
		log.Fatal(err)
		return err
	}
	fileInfo, _ := f.Stat()
	t.FilePath = fileInfo.Name()
	t.Title = m.Title()
	t.Album = m.Album()
	t.Artist = m.Artist()
	return nil
}
