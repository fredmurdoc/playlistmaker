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
			playlistmaker.LogInstance().Debug("call GetFirstEligibleTrack")
			t, err := playlistmaker.GetFirstEligibleTrack(path)
			if err != nil {
				playlistmaker.LogInstance().Warn(err.Error())
				t = nil
			}
			if t != nil {
				apiProvider := new(lastfm.LastFM)
				playlistmaker.LogInstance().Debug("call api.GetAlbumPlaylistFromAPIProviderByNameAndArtist")
				p := api.GetAlbumPlaylistFromAPIProviderByNameAndArtist(t, apiProvider)
				if p != nil {
					playlistmaker.LogInstance().Debug("call playlistmaker.FinalizeWithFilenames")
					errorFinalize := playlistmaker.FinalizeWithFilenames(p, path)
					if errorFinalize == nil {
						playlistmaker.LogInstance().Debug("call playlis.WriteToFile")
						isWrote := p.WriteToFileInDirectory(path)
						if isWrote {
							playlistmaker.LogInstance().Warn(fmt.Sprintf("Playlist created for %s", path))
						}
					} else {
						playlistmaker.LogInstance().Warn(errorFinalize.Error())
					}
				} else {
					playlistmaker.LogInstance().Warn(fmt.Sprintf("Found nothing for %s", path))
				}
			}
		}
	}

	//lastfm.GetAlbumMetadataFromMedia("my funny Valentine")
}
