package playlistmaker

import (
	"fmt"
)

type Playlist struct {
	Directory string
	Header    string
	Entries   []*PlaylistEntry
	Tail      string
	Filename  string
}

const firstLine string = "#EXTM3U"

//#EXTINF:durationinseconde, Sample artist - Sample title
var infoLineFormat string = "#EXTINF:%d, %s - %s\n%s"

//ToString:  retourne le contenu de la playlist sous forme de texte
func (p *Playlist) String() string {
	final := firstLine + "\n"
	for i := 0; i < len(p.Entries); i++ {
		final += p.Entries[i].String() + "\n"
	}
	return final
}

type PlaylistEntry struct {
	Track  *Track
	Order  int
	Length int
}

func (pe *PlaylistEntry) String() string {
	return fmt.Sprintf(infoLineFormat, pe.Length, pe.Track.Artist, pe.Track.Title, pe.Track.FilePath)
}

//Track track embedded in playlist
type Track struct {
	Title    string
	Album    string
	Artist   string
	FilePath string
	FileName string
}

//ToString get string representation
func (t *Track) String() string {
	return " - Title : " + t.Title +
		" - Artist : " + t.Artist +
		" - Album : " + t.Album
}
