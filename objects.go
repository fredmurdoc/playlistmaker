package playlistmaker

import (
	"fmt"
)

type Playlist struct {
	repertoire string
	header     string
	entries    []PlaylistEntry
	tail       string
	filename   string
}

const firstLine string = "#EXTM3U"

//#EXTINF:durationinseconde, Sample artist - Sample title
var infoLineFormat string = "#EXTINF:%d, %s - %s\n%s"

//ToString:  retourne le contenu de la playlist sous forme de texte
func (p *Playlist) String() string {
	final := firstLine + "\n"
	for i := 0; i < len(p.entries); i++ {
		final += p.entries[i].String() + "\n"
	}
	return final
}

type PlaylistEntry struct {
	track  Track
	order  int
	length int
}

func (pe *PlaylistEntry) String() string {
	return fmt.Sprintf(infoLineFormat, pe.length, pe.track.artist, pe.track.title, pe.track.filePath)
}

type Track struct {
	title    string
	album    string
	artist   string
	filePath string
}
