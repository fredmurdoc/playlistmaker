package tests

import (
	"log"
	"os"
	"strings"
	"testing"

	"github.com/fredmurdoc/playlistmaker"
)

//TestMetadataExtractor Test unit
func TestMetadataExtractor(m *testing.T) {
	file, errFile := os.OpenFile("../testdata/dirwithmultiplemedias/test02.mp3", os.O_RDONLY, os.FileMode(int(0755)))
	if errFile != nil {
		log.Fatal(errFile)
	}
	t := new(playlistmaker.Track)
	playlistmaker.ExtractMetadataToTrack(file, t)
	expected := "Duration :  - Title : Test02 - Artist : Test - Album : AlbumTest"

	got := t.String()
	if strings.Compare(expected, got) == 0 {
		m.Fatal("expected: '" + expected + "', got: '" + got + "'")
	}
}
