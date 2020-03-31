package playlistmaker

import (
	"os"
	"strings"
)

var playlistSuffix = []string{"m3u", "cue"}
var mediaSuffix = [3]string{"mp3", "flac", "ogg"}

type TrackFilter interface {
	isTrack(f os.FileInfo)
}

type PlaylistFilter interface {
	isPlaylist(f os.FileInfo)
}

//IsTrack : return true if file is a media music file
func IsTrack(f os.FileInfo) bool {
	for i := 0; i < len(mediaSuffix); i++ {
		if strings.HasSuffix(f.Name(), "."+mediaSuffix[i]) {
			return true
		}
	}
	return false
}

//IsPlaylist : return true if file is a playlist
func IsPlaylist(f os.FileInfo) bool {
	for i := 0; i < len(playlistSuffix); i++ {
		if strings.HasSuffix(f.Name(), "."+playlistSuffix[i]) {
			return true
		}
	}
	return false
}
