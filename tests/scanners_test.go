package tests

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/fredmurdoc/playlistmaker"
)

var baseDir = "../testdata"

func TestScannersFindSubDirectoriesWithNoPlaylistDirWithPlaylist(m *testing.T) {
	playlistmaker.FindSubDirectoriesWithNoPlaylist(baseDir)
	directoryWithPlaylist := baseDir + "/dirwithplaylist"
	if exists := playlistmaker.DirectoryWithNoPlaylist[directoryWithPlaylist]; exists {
		m.Fatal("Found a directory with a playlist" + directoryWithPlaylist)
	}

}

func TestScannersFindSubDirectoriesWithNoPlaylistDirWithNoMedia(m *testing.T) {
	playlistmaker.FindSubDirectoriesWithNoPlaylist(baseDir)
	directoryWithPlaylist := baseDir + "/dirwithoutmedia"
	if exists := playlistmaker.DirectoryWithNoPlaylist[directoryWithPlaylist]; exists {
		m.Fatal("Found a directory with no media" + directoryWithPlaylist)
	}

}

func TestScannersFindSubDirectoriesWithNoPlaylistDirWithOutPlaylist(m *testing.T) {
	playlistmaker.FindSubDirectoriesWithNoPlaylist(baseDir)
	directoryWithNoPlaylist := baseDir + "/dirwithoutplaylist"
	if exists := playlistmaker.DirectoryWithNoPlaylist[directoryWithNoPlaylist]; !exists {
		m.Fatal("No Found a directory with no playlist" + directoryWithNoPlaylist)
	}
}

func TestScannersGetFirstEligibleTrack(m *testing.T) {
	t, _ := playlistmaker.GetFirstEligibleTrack(baseDir + "/dirwithmultiplemedias")
	expected := "test01.ogg"
	if t.FileName != expected {
		m.Fatal("No Found expected : " + expected + ", got: " + t.FileName)
	}
}

//dirwithmixedcontent
func TestScannersDistanceWithTracks(m *testing.T) {
	tested := "Track Test In Name"

	filepath.Walk(baseDir+"/dirfordistance", func(path string, f os.FileInfo, err error) error {
		if playlistmaker.IsTrack(f) {
			name := strings.TrimSuffix(filepath.Base(path), filepath.Ext(path))
			playlistmaker.IsNameMatch(name, tested)
			playlistmaker.IsNameMatch(name, "01 "+tested)
		}
		return nil
	})
}

func TestScannersFinalizeWithFilenames(m *testing.T) {

	playlistmaker.LogInstance().SetLevel(playlistmaker.Debug)
	playlistmaker.LogInstance().Info("TestScannersFinalizeWithFilenames")
	//notEligible := [1]string{"track not eligible for distance.ogg"}
	eligibles := []string{"01 - TrackTestInName.ogg", "TrackTestInName.ogg", "track-test-in-name.ogg", "01 - track-test-in-name.ogg"}

	p := playlistmaker.Playlist{}
	for index, item := range eligibles {
		pe := playlistmaker.PlaylistEntry{}
		pe.Order = index
		pe.Length = 100 + index
		trackName := strings.TrimSuffix(filepath.Base(item), filepath.Ext(item))
		pe.Track = &playlistmaker.Track{Title: trackName, Album: "AlbumTest", Artist: "ArtistTest"}

		p.Entries = append(p.Entries, &pe)
	}
	errorFinalize := playlistmaker.FinalizeWithFilenames(&p, baseDir+"/dirfordistance")
	if errorFinalize == nil {
		fmt.Println(p.String())
		isCompleted := p.IsCompleted()
		if !isCompleted {
			m.Fatal("Playlist is not Completed !!")
		}
		for _, itemExpected := range eligibles {
			found := false
			for _, itemGot := range p.Entries {
				found = found || (itemGot.Track.FileName == itemExpected)
			}
			if !found {
				m.Fatal(itemExpected + " is not in Playlist  !!")
			}
		}
	} else {

		playlistmaker.LogInstance().Warn(errorFinalize.Error())
	}
}
