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
	//DirectoryWithAPIFailure scanned directory playlist failed to make
	DirectoryWithAPIFailure map[string]bool
	//DirectoryWithPlaylistFailure scanned directory playlist failed to make
	DirectoryWithPlaylistFailure map[string]bool

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
	DirectoryWithAPIFailure = make(map[string]bool)
	DirectoryWithPlaylistFailure = make(map[string]bool)
	filepath.Walk(root, DirectoriesWithoutPlaylistVisitor)
}

//isTrackIsEligibleForAPICall : determine si les informations du Track sont suffisante, retourne vrai ou faux{
func isTrackIsEligibleForAPICall(t *Track) bool {
	LogInstance().Debug("isTrackIsEligibleForAPICall : " + t.String())
	return t.Artist != "" && t.Album != "" && t.Title != ""
}

//GetFirstEligibleTrack : prend le premier fichier musical eligible, retourne Track eligible
func GetFirstEligibleTrack(repertoire string) (*Track, error) {
	var eligible *Track
	eligible = nil
	errWalk := filepath.Walk(repertoire, func(path string, f os.FileInfo, err error) error {
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
				return errExtract
			}
			if isTrackIsEligibleForAPICall(t) {
				LogInstance().Info("is eligible : " + path)
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
	return eligible, errWalk
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
		LogInstance().Info("findCorrespondingEntryInPlaylist for " + pathMedia)
		err = findCorrespondingEntryInPlaylist(pathMedia, p)
	}
	return err
}

func findCorrespondingEntryInPlaylist(mediaFilename string, p *Playlist) error {

	t := new(Track)
	f, errOpen := os.Open(mediaFilename)
	if errOpen != nil {
		LogInstance().Debug("error opening : " + mediaFilename)
		log.Fatal(errOpen)
		return errOpen
	}
	LogInstance().Debug("extract metadata : " + mediaFilename)
	errExtract := ExtractMetadataToTrack(f, t)
	LogInstance().Debug("extract metadata error return : ")
	LogInstance().Debug(errExtract)
	LogInstance().Debug("extract metadata Track : " + t.String())
	if errExtract != nil {
		LogInstance().Debug("extract metadat error on : " + mediaFilename)
	}

	for _, entry := range p.Entries {
		if t.Title != "" {
			LogInstance().Debug(fmt.Sprintf("try with extracted Track %s", t.String()))
			findBestMatch(t.Title, entry, t.FileName)
		} else {
			relativeMediaName := strings.TrimSuffix(filepath.Base(mediaFilename), filepath.Ext(mediaFilename))
			LogInstance().Debug(fmt.Sprintf("try with filename %s", relativeMediaName))
			findBestMatch(relativeMediaName, entry, mediaFilename)
		}

		LogInstance().Debug("entry founded :  " + entry.String())
	}
	//control playlist
	filesInPlaylist := make(map[string]bool)
	for _, entry := range p.Entries {
		currentFilename := entry.Track.FileName
		if len(currentFilename) > 0 {
			if _, exists := filesInPlaylist[currentFilename]; exists {
				LogInstance().Warn(fmt.Sprintf("%s is already in control list", currentFilename))
				return errors.New("Playlist with multiple entries with same file")
			}
			LogInstance().Debug(fmt.Sprintf("append %s to control list", currentFilename))
			filesInPlaylist[currentFilename] = true
		}
	}

	return nil
}

func findBestMatch(searched string, entry *PlaylistEntry, filename string) bool {
	var (
		hasFound    bool
		actualScore float64
		newScore    float64
	)
	//if we have already found a filaneme, so we have a BestScore
	if entry.Track.FileName != "" {
		actualScore = entry.BestScore
		LogInstance().Debug(fmt.Sprintf("Set actual score %.6f to existing filename in track %s <-> %s", actualScore, entry.Track.FileName, entry.Track.Title))
	}

	newScore = getScore(searched, entry.Track.Title)

	if newScore > actualScore {
		LogInstance().Debug(fmt.Sprintf("new score %.6f with title %s <-> %s", newScore, searched, entry.Track.Title))
		LogInstance().Debug(fmt.Sprintf("replace filename %s with  %s", entry.Track.FileName, filepath.Base(filename)))
		entry.Track.FileName = filepath.Base(filename)
		entry.BestScore = newScore
		actualScore = newScore
		hasFound = true
	} else {
		entry.BestScore = actualScore
		LogInstance().Debug(fmt.Sprintf("no new score %.6f with title %s <-> %s", newScore, searched, entry.Track.Title))
	}
	if entry.Order > 0 {
		newScore = getScore(searched, strconv.Itoa(entry.Order)+entry.Track.Title)

		if newScore > actualScore {
			LogInstance().Debug(fmt.Sprintf("new score %.6f with title %s <-> %s", newScore, searched, strconv.Itoa(entry.Order)+entry.Track.Title))
			LogInstance().Debug(fmt.Sprintf("replace filename %s with  %s", entry.Track.FileName, filepath.Base(filename)))
			entry.Track.FileName = filepath.Base(filename)
			entry.BestScore = newScore
			actualScore = newScore
			hasFound = true
		} else {
			entry.BestScore = actualScore
			LogInstance().Debug(fmt.Sprintf("no new score %.6f with title %s <-> %s", newScore, searched, strconv.Itoa(entry.Order)+entry.Track.Title))
		}
	}
	return hasFound
}

func getScore(tested string, target string) float64 {
	return levenshtein.RatioForStrings([]rune(tested), []rune(target), DistanceDefaultOptionsWithSub)
}

//IsNameMatch : return if name match tarcget
func IsNameMatch(tested string, target string) bool {
	distance := getScore(tested, target)
	LogInstance().Debug("tested: " + tested + ", target: " + target + " : distance is " + fmt.Sprint("%.6f", distance))
	return distance >= threasholdRatioForDistance
}
