# Conception technique

## Synopsis

Le programme scanne chaque répertoire et détermine ceux qui ne possèdent pas de playlist.

Pour chaque répertoire qui ne possède pas de playlist, le programme prend le premier morceau, et en extrait l'artiste et le nom de l'album.
Puis le programme utilise l'API de recherche du contenu de l'album et pour constituer la playlist.

 - s'il ne trouve pas ces informations, le programme abandonne la création de la playlist pour ce répertoire.
Pour chaque titre de l'album issu de l'API le programme fait une correspondance avec les fichiers musicaux. La correspondance doit être complète.

 - si la correpondance n'est pas complète, le programme abandonne la création de la playlist pour ce répertoire.

Un fichier de playlist mentionne pour chaque titre

 - l'ordre du morceau dans l'album
 - la durée en seconde du titre 
 - le nom du morceau
 - le fichier correspondant (relatif au répertoire)

Une playlist peut contenir des données techniques lié à son format en entete (_header_) ou en fin de fichier (_tail_)

 ## Conception du code

```
 Objet Playlist {
     repertoire
     header
     PlaylistEntry[]
     tail
     filename
     ToString() :  retourne le contenu de la playlist sous forme de texte
 }
 

 Objet PlaylistEntry {
     Track
     order
     length
 }

Objet Track {
    title 
    album
    artist
    filePath
}

package filters {
    TrackFilter{
        isTrack(file)
    }

    PlaylistFilter{
        isPlaylist(file)
    }
}

TrackMetadataExtractor{
    ExtractToTrack(file, Track) : alimente le Track avec les metadonnées extraites dans le fichier de media
}
package scanners {
    PlaylistScanner {
        //FindDirectoriesWithoutPlaylist :  retourne list de repertoires qui ne possèdent pas de playlist
        FindDirectoriesWithoutPlaylist(repertoireRacine)
    }

    //EligibleApiCallTrackScanner :  se charge de trouver le premier track eligible à un apel API sur la base d'un repertoire
    EligibleApiCallTrackScanner{
        TrackFilter
        TrackMetadataExtractor
        isTrackIsEligibleForApiCall(Track) : determine si les informations du Track sont suffisante, retourne vrai ou faux{
            return Track.artist != '' &&
            Track.album != '' &&
            Track.title != ''
        }
        //getFirstEligibleTrack: prend le premier fichier musical eligible, retourne Track eligible
        getFirstEligibleTrack(repertoire)  {
            foreach(file in repertoire){
                if(TrackFilter.isTrack(file)){
                    t = new Track()
                    TrackMetadataExtractor.ExtractToTrack(file, t)
                    if isTrackIsEligibleForApiCall(t){
                        return t
                    }
                }
            }
        }    
    }
}
PlaylistWriter {
    //Write : écrit le fichier playlist dans le format attendu
    Write(Playlist)
}

package api {

    PlaylistApiResult{
        album
        artist
        ordre
        duree
    }
    PlaylistApiInterface{
    //getApiResult: appelle l'api et retourne les resultats
        getApiResult(Track){
            result = callAlbumAPI(Track.album, Track.artist)
        }
    }

    PlaylistApi {
        PlaylistApiInterface
        //GetAlbumPlaylistFromNameAndArtist: retourne PlaylistEntry[] ou rien
        GetAlbumPlaylistFromNameAndArtist(Track){
            PlaylistApiResult[] results = PlaylistApiInterface.getApiResult(Track.album, Track.artist)
            entries = getPlaylistEntriesFromApiResults(results)
            Playlist = new Playlist(entries)
            return Playlist    
        }
        

        //getPlaylistEntriesFromApiResults:  parse les resultats de l'API 
        getPlaylistEntriesFromApiResults(results){
                entries []
                foreach(result in results){
                    PlaylistEntry.Track = new Track
                    PlaylistEntry.Track.title 
                    PlaylistEntry.Track.album = result.album
                    PlaylistEntry.Track.artist = result.artist
                    PlaylistEntry.order = result.ordre
                    PlaylistEntry.length = result.duree
                    PlaylistEntry.Track.filePath = null
                    entries.append(PlaylistEntry)
                }
                return entries
            }
    }
}
Program{
    PlaylistScanner
    PlaylistFabric
    PlaylistApi
    main(){
        PlaylistScanner.FindDirectoriesWithoutPlaylist(repertoireRacine) : path[]
        foreach(path in path[]){
            Track = PlaylistFabric.GetFirstEligibleTrack(repertoire)
            Playlist = PlaylistApi.GetAlbumPlaylistFromNameAndArtist(Track)    
            PlaylistFabric.Write(Playlist)
        }
    }
}

```