package tests

import (
	"testing"

	"github.com/fredmurdoc/playlistmaker"
)

var baseDir = "../testdata"

func TestFindSubDirectoriesWithNoPlaylistDirWithPlaylist(m *testing.T) {
	playlistmaker.FindSubDirectoriesWithNoPlaylist(baseDir)
	directoryWithPlaylist := baseDir + "/dirwithplaylist"
	if exists := playlistmaker.DirectoryWithNoPlaylist[directoryWithPlaylist]; exists {
		m.Fatal("Found a directory with a playlist" + directoryWithPlaylist)
	}

}

func TestFindSubDirectoriesWithNoPlaylistDirWithNoMedia(m *testing.T) {
	playlistmaker.FindSubDirectoriesWithNoPlaylist(baseDir)
	directoryWithPlaylist := baseDir + "/dirwithoutmedia"
	if exists := playlistmaker.DirectoryWithNoPlaylist[directoryWithPlaylist]; exists {
		m.Fatal("Found a directory with no media" + directoryWithPlaylist)
	}

}

func TestFindSubDirectoriesWithNoPlaylistDirWithOutPlaylist(m *testing.T) {
	playlistmaker.FindSubDirectoriesWithNoPlaylist(baseDir)
	directoryWithNoPlaylist := baseDir + "/dirwithoutplaylist"
	if exists := playlistmaker.DirectoryWithNoPlaylist[directoryWithNoPlaylist]; !exists {
		m.Fatal("No Found a directory with no playlist" + directoryWithNoPlaylist)
	}
}

//dirwithmixedcontent
