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
			apiProvider := new(lastfm.LastFM)
			p := api.GetAlbumPlaylistFromAPIProviderByNameAndArtist(t, apiProvider)
			playlistmaker.FinalizeWithFilenames(p, path)
			fmt.Println("Result : " + p.String())
		}
	}

	//lastfm.GetAlbumMetadataFromMedia("my funny Valentine")
}
