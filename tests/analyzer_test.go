package medias

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/fredmurdoc/playlistmaker/medias"
)

//TestMetadataExtractor Test unit
func TestMetadataExtractor(m *testing.T) {
	file, errFile := os.OpenFile("/home/fred/Musique/Steve Coleman - Functional Arrhythmias (2013) [EAC-FLAC]/02 - Medulla-Vagus.flac", os.O_RDONLY, os.FileMode(int(0755)))
	if errFile != nil {
		log.Fatal(errFile)
	}
	media, _ := medias.ExtractMetadata(file)
	fmt.Println("Result : " + media.ToString())
}