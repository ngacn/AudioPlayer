package player

import (
	"warpten/playlists"
	"warpten/tracks"
)

var version string
var pls playlists.Playlists
var tks tracks.Tracks

func Version() string {
	return version
}

func Playlists() playlists.Playlists {
	return pls
}

func Tracks() tracks.Tracks {
	return tks
}

func Playlist(name string) ([]string, bool) {
	return pls.Playlist(name)
}

func AddPlaylist(name string) error {
	return pls.AddPlaylist(name)
}

func DelPlaylist(name string) error {
	uuids, exists := pls.Playlist(name)
	if !exists {
		return playlists.ErrPlaylistNotExists
	}

	for _, uuid := range uuids {
		if err := tks.DelTrack(uuid); err != nil {
			return err
		}
	}

	if err := pls.DelPlaylist(name); err != nil {
		return err
	}
	return nil
}

func Track(uuid string) (*tracks.Track, bool) {
	tk, exists := tks.Track(uuid)
	return tk, exists
}

func AddTrack(path, playlist string) error {
	uuid, err := tks.AddTrack(path)
	if err != nil {
		return err
	}

	_, exists := pls.Playlist(playlist)
	if exists {
		pls.AddUUIDs(playlist, uuid)
		return nil
	}
	return playlists.ErrPlaylistNotExists
}

func DelTrack(uuid, playlist string) error {
	_, exists := pls.Playlist(playlist)
	if exists {
		pls.DelUUIDs(playlist, uuid)
	} else {
		return playlists.ErrPlaylistNotExists
	}
	return tks.DelTrack(uuid)
}

func Init() {
	version = "0.0"
	pls = playlists.New()
	pls.AddPlaylist("Default")
	tks = tracks.New()
}
