package medias

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

//MusicFilesList variable with all music files
var MusicFilesList []string

//DirectoriesWithPlaylist variable with all direcotries with playlist files
var DirectoriesWithPlaylist []string

//FindMissingMusicFilesList get all files not associated with a playlist
func FindMissingMusicFilesList() []string {
	missingMusicFilesList := []string{}
	for i := 0; i < len(MusicFilesList); i++ {
		for j := 0; j < len(DirectoriesWithPlaylist); j++ {
			if !strings.HasPrefix(MusicFilesList[i], DirectoriesWithPlaylist[j]) {
				missingMusicFilesList = append(missingMusicFilesList, MusicFilesList[i])
			}
		}
	}
	return missingMusicFilesList
}

//Visitor visitor function for path/filepath.Walk for find music files
func Visitor(path string, f os.FileInfo, err error) error {
	fmt.Printf("Visited: %s\n", path)
	playlistSuffix := []string{"m3u"}
	searchedSuffix := [3]string{"mp3", "flac", "ogg"}
	for i := 0; i < len(searchedSuffix); i++ {
		if strings.HasSuffix(f.Name(), "."+searchedSuffix[i]) {
			if len(MusicFilesList) == 0 {
				MusicFilesList = []string{path}
			} else {
				MusicFilesList = append(MusicFilesList, path)
			}
		}
	}
	for i := 0; i < len(playlistSuffix); i++ {
		if strings.HasSuffix(f.Name(), "."+playlistSuffix[i]) {
			parentDir, errGetDir := filepath.Abs(filepath.Dir(path))
			if errGetDir != nil {
				log.Fatal(errGetDir)
			}
			DirectoriesWithPlaylist = append(DirectoriesWithPlaylist, parentDir)
		}
	}
	return nil
}
