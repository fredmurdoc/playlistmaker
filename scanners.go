package playlistmaker

import (
	"os"
	"path/filepath"
)

var DirectoryWithAtLeastMediaFile map[string]bool
var DirectoryWithNoPlaylist map[string]bool
var DirectoryWithPlaylist map[string]bool

type PlaylistScanner struct {
}

//DirectoriesWithoutPlaylistVisitor: visitor for walk
func DirectoriesWithoutPlaylistVisitor(path string, f os.FileInfo, err error) error {
	if IsTrack(f) {
		DirectoryWithAtLeastMediaFile[filepath.Dir(path)] = true
	}
	if IsPlaylist(f) {
		DirectoryWithPlaylist[filepath.Dir(path)] = true
		DirectoryWithAtLeastMediaFile[filepath.Dir(path)] = true
	} else {
		//si on est dans un repertoire dans lequel on a un fichier media ou playlist
		if DirectoryWithAtLeastMediaFile[filepath.Dir(path)] {
			//si on a pas déja deja croisé un fichier playlist
			if !DirectoryWithPlaylist[filepath.Dir(path)] {
				DirectoryWithNoPlaylist[filepath.Dir(path)] = true
			} else {
				//dans le cas contraire on le marque à false
				DirectoryWithNoPlaylist[filepath.Dir(path)] = false
			}
		}
	}
	return nil
}

//FindSubDirectoriesWithNoPlaylist :  retourne list de repertoires qui ne possèdent pas de playlist
func FindSubDirectoriesWithNoPlaylist(root string) {
	DirectoryWithAtLeastMediaFile = make(map[string]bool)
	DirectoryWithPlaylist = make(map[string]bool)
	DirectoryWithNoPlaylist = make(map[string]bool)
	filepath.Walk(root, DirectoriesWithoutPlaylistVisitor)
}
