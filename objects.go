package playlistmaker

import (
	"fmt"
)

// -----
type Playlist struct {
	Directory string
	Header    string
	Entries   []*PlaylistEntry
	Tail      string
	FileName  string
}

const firstLine string = "#EXTM3U"

//#EXTINF:durationinseconde, Sample artist - Sample title
var infoLineFormat string = "#EXTINF:%d, %s - %s\n%s"

//ToString:  retourne le contenu de la playlist sous forme de texte
func (p *Playlist) String() string {
	final := firstLine + "\n"
	for i := 0; i < len(p.Entries); i++ {
		final += p.Entries[i].String() + "\n\n"
	}
	return final
}

//IsCompleted return if Playlist is complete for playlist file
func (p *Playlist) IsCompleted() bool {
	result := true
	for _, entry := range p.Entries {
		result = result && entry.IsCompleted()
	}
	return result
}

// -----

type PlaylistEntry struct {
	Track  *Track
	Order  int
	Length int
}

func (pe *PlaylistEntry) String() string {
	return fmt.Sprintf(infoLineFormat, pe.Length, pe.Track.Artist, pe.Track.Title, pe.Track.FileName)
}

//IsCompleted return if PlaylistEntry is complete for playlist file
func (pe *PlaylistEntry) IsCompleted() bool {
	return pe.Length != 0 && pe.Track.Artist != "" && pe.Track.Title != "" && pe.Track.FileName != ""
}

// -----

//Track track embedded in playlist
type Track struct {
	Title    string
	Album    string
	Artist   string
	FileName string
}

//ToString get string representation
func (t *Track) String() string {
	return " - Title : " + t.Title +
		" - Artist : " + t.Artist +
		" - Album : " + t.Album
}
