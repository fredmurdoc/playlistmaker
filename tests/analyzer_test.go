package tests

import (
	"log"
	"os"
	"strings"
	"testing"

	"github.com/fredmurdoc/playlistmaker/medias"
)

//TestMetadataExtractor Test unit
func TestMetadataExtractor(m *testing.T) {
	file, errFile := os.OpenFile("../testdata/dirwithmultiplemedias/test02.mp3", os.O_RDONLY, os.FileMode(int(0755)))
	if errFile != nil {
		log.Fatal(errFile)
	}
	media, _ := medias.ExtractMetadata(file)
	expected := "Duration :  - Title : Test02 - Artist : Test - Album : AlbumTest"

	got := media.ToString()
	if strings.Compare(expected, got) == 0 {
		m.Fatal("expected: '" + expected + "', got: '" + got + "'")
	}
}
