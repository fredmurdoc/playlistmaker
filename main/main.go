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
	fmt.Println("version 0.3")
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

	if len(playlistmaker.DirectoryWithPlaylist) > 0 {
		fmt.Println("---------------------------------------------")
		fmt.Println("Directories scanned already with playlist")
		fmt.Println("---------------------------------------------")
		for entry, exists := range playlistmaker.DirectoryWithPlaylist {
			if exists {
				fmt.Println(entry)
			}
		}
		fmt.Println("")
	}
	if len(playlistmaker.DirectoryWithAPIFailure) > 0 {
		fmt.Println("---------------------------------------------")
		fmt.Println("Directories scanned but fail to find playlist")
		fmt.Println("---------------------------------------------")
		for entry, exists := range playlistmaker.DirectoryWithAPIFailure {
			if exists {
				fmt.Println(entry)
			}
		}
		fmt.Println("")
	}
	if len(playlistmaker.DirectoryWithPlaylistFailure) > 0 {
		fmt.Println("-----------------------------------------------")
		fmt.Println("Directories scanned but fail to create playlist")
		fmt.Println("-----------------------------------------------")
		for entry, exists := range playlistmaker.DirectoryWithPlaylistFailure {
			if exists {
				fmt.Println(entry)
			}
		}
	}
}
