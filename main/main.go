package main

import (
	//"github.com/fredmurdoc/playlistmaker/lastfm"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/fredmurdoc/playlistmaker/medias"
)

func main() {
	flag.Parse()
	root := flag.Arg(0)
	err := filepath.Walk(root, medias.Visitor)
	fmt.Printf("filepath.Walk() returned %v\n", err)
	list := medias.FindMissingMusicFilesList()
	for i := 0; i < len(list); i++ {
		fmt.Printf("file : " + list[i] + "\n")
		musicFile := list[i]

		f, errFile := os.Open(musicFile)
		if errFile != nil {
			log.Fatalln("ERROR ON : " + musicFile)
			log.Fatal(errFile)
		}
		media, _ := medias.ExtractMetadata(f)
		fmt.Println("Result : " + media.ToString())
	}

	//lastfm.GetAlbumMetadataFromMedia("my funny Valentine")
}
