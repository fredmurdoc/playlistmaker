package playlistmaker

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
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

func (p *Playlist) Len() int           { return len(p.Entries) }
func (p *Playlist) Swap(i, j int)      { p.Entries[i], p.Entries[j] = p.Entries[j], p.Entries[i] }
func (p *Playlist) Less(i, j int) bool { return p.Entries[i].Order < p.Entries[j].Order }

//ToString:  retourne le contenu de la playlist sous forme de texte
func (p *Playlist) String() string {
	final := firstLine + "\n"
	sort.Sort(p)
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

//WriteToFileInDirectory write playlist content in directory
func (p *Playlist) WriteToFileInDirectory(dirName string) bool {
	fileName := p.Entries[0].Track.Album + ".m3u"
	fileName = strings.Replace(fileName, string(os.PathSeparator), "-", -1)
	finalFileName := filepath.FromSlash(dirName + "/" + fileName)
	fd, errCreate := os.Create(finalFileName)
	if errCreate != nil {
		log.Fatalln(errCreate)
		fd.Close()
	}
	nb, errWrite := fd.WriteString(p.String())
	if errWrite != nil {
		log.Fatalln(errWrite)
		fd.Close()
		return false
	}
	fd.Close()
	return nb > 0

}

// -----

type PlaylistEntry struct {
	Track     *Track
	Order     int
	Length    int
	BestScore float64
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
