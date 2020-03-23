package medias

import (
	"fmt"
)

const firstLine string = "#EXTM3U"

//Playlist type
type Playlist struct {
	//Tracks
	Tracks []string
}

//#EXTINF:durationinseconde, Sample artist - Sample title
var infoLineFormat string = "#EXTINF:%d, %s - %s\n%s"

//AppendTrack to playlist
func (p *Playlist) AppendTrack(track *TrackMetaData) {
	p.Tracks = append(p.Tracks, fmt.Sprintf(infoLineFormat, track.DurationInSecond, track.Artist, track.Name, track.Filepath))
}

//ToString retrun string content of playlist
func (p *Playlist) ToString() string {
	final := firstLine + "\n"
	for i := 0; i < len(p.Tracks); i++ {
		final += p.Tracks[i] + "\n"
	}
	return final
}

//Length length of playlist
func (p *Playlist) Length() int {
	return len(p.Tracks)
}
