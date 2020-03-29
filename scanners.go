package playlistmaker

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/texttheater/golang-levenshtein/levenshtein"
)

const threasholdRatioForDistance = 0.8

var (
	//DirectoryWithAtLeastMediaFile scanned directory with at lest one media file inside
	DirectoryWithAtLeastMediaFile map[string]bool
	//DirectoryWithNoPlaylist scanned directory with no playlist file inside
	DirectoryWithNoPlaylist map[string]bool
	//DirectoryWithPlaylist scanned directory with at lest one playlist file inside
	DirectoryWithPlaylist map[string]bool
	//DistanceDefaultOptionsWithSub options for levenshtein distance calculations
	DistanceDefaultOptionsWithSub levenshtein.Options = levenshtein.Options{
		InsCost: 1,
		DelCost: 1,
		SubCost: 1,
		Matches: levenshtein.IdenticalRunes,
	}
)

//DirectoriesWithoutPlaylistVisitor visitor for walk
func DirectoriesWithoutPlaylistVisitor(path string, f os.FileInfo, err error) error {

	//Si playlist
	if IsPlaylist(f) {
		DirectoryWithNoPlaylist[filepath.Dir(path)] = false
		DirectoryWithPlaylist[filepath.Dir(path)] = true
		DirectoryWithAtLeastMediaFile[filepath.Dir(path)] = true
		LogInstance().Debug(path + " is a playlist")
	} else { //Si autre chose
		if IsTrack(f) {
			DirectoryWithAtLeastMediaFile[filepath.Dir(path)] = true
			LogInstance().Debug(path + " is a track")

			//si on est dans un repertoire dans lequel on a un fichier media ou playlist
			if DirectoryWithAtLeastMediaFile[filepath.Dir(path)] {
				LogInstance().Debug(filepath.Dir(path) + " a au moins un fichier playlist ou media")
				//si on a pas déja deja croisé un fichier playlist
				exists := DirectoryWithPlaylist[filepath.Dir(path)]
				if !exists {
					LogInstance().Debug(filepath.Dir(path) + " n'a pas  de fichier playlist")
					DirectoryWithNoPlaylist[filepath.Dir(path)] = true

				} else {
					LogInstance().Debug(filepath.Dir(path) + " a un fichier playlist")
					//dans le cas contraire on le marque à false
					DirectoryWithNoPlaylist[filepath.Dir(path)] = false
				}
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

//isTrackIsEligibleForAPICall : determine si les informations du Track sont suffisante, retourne vrai ou faux{
func isTrackIsEligibleForAPICall(t *Track) bool {
	LogInstance().Debug("testTrack : " + t.String())
	return t.Artist != "" && t.Album != "" && t.Title != ""
}

//GetFirstEligibleTrack : prend le premier fichier musical eligible, retourne Track eligible
func GetFirstEligibleTrack(repertoire string) *Track {
	var eligible *Track
	eligible = nil
	filepath.Walk(repertoire, func(path string, f os.FileInfo, err error) error {
		if eligible != nil {
			return nil
		}
		LogInstance().Debug("scan : " + path)
		if IsTrack(f) {
			LogInstance().Debug("is track : " + path)
			t := new(Track)
			f, err := os.Open(path)
			if err != nil {
				LogInstance().Debug("error opening : " + path)
				log.Fatal(err)
				return err
			}
			LogInstance().Debug("extract metadata : " + path)
			errExtract := ExtractMetadataToTrack(f, t)
			LogInstance().Debug("extract metadata return : ")
			LogInstance().Debug(errExtract)
			if errExtract != nil {
				LogInstance().Debug("error on : " + path)
				log.Fatalln(errExtract)
				return errExtract
			}
			if isTrackIsEligibleForAPICall(t) {
				LogInstance().Debug("is eligible : " + path)
				eligible = t
				return nil
			} else {
				LogInstance().Debug("not eligible : " + path)
			}
		} else {
			LogInstance().Debug("is not a track : " + path)
		}
		return nil
	})
	return eligible
}

//FinalizeWithFilenames finalize playlist items with filenames in directory
func FinalizeWithFilenames(p *Playlist, directory string) error {
	var err error
	mediasList := make(map[string]os.FileInfo)
	filepath.Walk(directory, func(path string, f os.FileInfo, err error) error {
		if IsTrack(f) {
			mediasList[path] = f
		}
		return nil
	})
	for pathMedia := range mediasList {
		ok := findCorrespondingEntryInPlaylist(pathMedia, p)
		if !ok {
			err = errors.New("cannot find playlist entry for media " + pathMedia)
		}
	}
	return err
}

func findCorrespondingEntryInPlaylist(mediaFilename string, p *Playlist) bool {
	relativeMediaName := strings.TrimSuffix(filepath.Base(mediaFilename), filepath.Ext(mediaFilename))
	hasFound := false
	for _, entry := range p.Entries {
		if IsNameMatch(relativeMediaName, entry.Track.Title) {
			entry.Track.FileName = filepath.Base(mediaFilename)
			hasFound = true
		} else {
			if IsNameMatch(relativeMediaName, strconv.Itoa(entry.Order)+entry.Track.Title) {
				entry.Track.FileName = filepath.Base(mediaFilename)
				hasFound = true
			}
		}
	}
	return hasFound
}

//IsNameMatch : return if name match tarcget
func IsNameMatch(tested string, target string) bool {

	distance := levenshtein.RatioForStrings([]rune(tested), []rune(target), DistanceDefaultOptionsWithSub)
	LogInstance().Debug("tested: " + tested + ", target: " + target + " : distance is " + fmt.Sprint("%d", distance))
	return distance >= threasholdRatioForDistance
}
