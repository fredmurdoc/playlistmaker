package main

import (
	"flag"
	"fmt"

	"github.com/fredmurdoc/playlistmaker"
	"github.com/fredmurdoc/playlistmaker/api"
	"github.com/fredmurdoc/playlistmaker/api/lastfm"
)

func main() {
	flag.Parse()
	root := flag.Arg(0)
	playlistmaker.LogInstance().SetLevel(playlistmaker.Debug)
	playlistmaker.FindSubDirectoriesWithNoPlaylist(root)

	for path, exists := range playlistmaker.DirectoryWithNoPlaylist {
		if exists {
			t := playlistmaker.GetFirstEligibleTrack(path)
			if t != nil {
				apiProvider := new(lastfm.LastFM)
				p := api.GetAlbumPlaylistFromAPIProviderByNameAndArtist(t, apiProvider)
				if p != nil {
					playlistmaker.FinalizeWithFilenames(p, path)
					isWrote := p.WriteToFile(path + "/" + p.Entries[0].Track.Album + ".m3u")
					if isWrote {
						playlistmaker.LogInstance().Warn(fmt.Sprintf("Playlist created for %s", path))
					}
				} else {
					playlistmaker.LogInstance().Warn(fmt.Sprintf("Found nothing for %s", path))
				}
			}
		}
	}

	//lastfm.GetAlbumMetadataFromMedia("my funny Valentine")
}
