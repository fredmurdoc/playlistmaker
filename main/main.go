package main

import (
	"flag"
	"fmt"

	"github.com/fredmurdoc/playlistmaker"
	"github.com/fredmurdoc/playlistmaker/api"
	"github.com/fredmurdoc/playlistmaker/api/lastfm"
)

func main() {

	var root, loglevel string

	flag.StringVar(&root, "directory", ".", "directory to scan")
	flag.StringVar(&loglevel, "loglevel", "warn", "loglevel : debug, info, warn")
	fmt.Println("version 0.4")
	flag.Parse()

	switch {
	case loglevel == "debug":
		playlistmaker.LogInstance().SetLevel(playlistmaker.Debug)
	case loglevel == "info":
		playlistmaker.LogInstance().SetLevel(playlistmaker.Info)
	case loglevel == "warn":
		playlistmaker.LogInstance().SetLevel(playlistmaker.Warn)
	}

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
							playlistmaker.DirectoryWithPlaylistSuccess[path] = true
						}
					} else {
						playlistmaker.LogInstance().Warn(errorFinalize.Error())
						playlistmaker.DirectoryWithPlaylistFailure[path] = true
					}
				} else {
					playlistmaker.LogInstance().Warn(fmt.Sprintf("Found nothing for %s", path))
					playlistmaker.DirectoryWithAPIFailure[path] = true
				}
			}
		}
	}

	report(playlistmaker.DirectoryWithPlaylist, "Directories scanned already with playlist")
	report(playlistmaker.DirectoryWithAPIFailure, "Directories scanned but fail to find playlist")
	report(playlistmaker.DirectoryWithPlaylistFailure, "Directories scanned but fail to create playlist")
	report(playlistmaker.DirectoryWithPlaylistSuccess, "Directories scanned with new playlist")

}

func report(list map[string]bool, message string) {
	if len(list) > 0 {
		fmt.Println("-----------------------------------------------")
		fmt.Println(message)
		fmt.Println("-----------------------------------------------")
		for entry, exists := range list {
			if exists {
				fmt.Println(entry)
			}
		}
		fmt.Println("")
	}
}
