package player

import (
	"errors"
	"os/exec"
	"runtime"
	"warpten/playlist"
	"warpten/track"
)

var (
	ErrPlaylistExists    = errors.New("Playlist already exists")
	ErrPlaylistNotExists = errors.New("Playlist not exists")

	ErrNotImplemented = errors.New("Method not implemented")
)

var version string
var playlists map[string]*playlist.Playlist
var tracks map[string]*track.Track

func Version() string {
	return version
}

func Playlists() map[string]*playlist.Playlist {
	return playlists
}

func Playlist(name string) (*playlist.Playlist, bool) {
	pl, exists := playlists[name]
	return pl, exists
}

func NewPlaylist(name string) error {
	if _, exists := playlists[name]; exists {
		return ErrPlaylistExists
	}
	playlists[name] = playlist.New()
	return nil
}

func DelPlaylist(name string) error {
	if pl, exists := playlists[name]; exists {
		pl.Clear()
		delete(playlists, name)
		return nil
	} else {
		return ErrPlaylistNotExists
	}
}

func NewTrack(path, playlist string) error {
	var uuid string
	switch os := runtime.GOOS; os {
	case "darwin":
	case "linux":
		for {
			out, err := exec.Command("uuidgen").Output()
			if err != nil {
				return err
			}
			uuid = string(out)
			if _, exists := tracks[uuid]; !exists {
				break
			}
		}
	default:
		return ErrNotImplemented
	}

	t := track.New(uuid, path)
	tracks[uuid] = t
	pl, exists := Playlist("Default")
	if exists {
		pl.Append(t)
	} else {
		return ErrPlaylistNotExists
	}
	return nil
}

func Track(uuid string) (*track.Track, bool) {
	t, exists := tracks[uuid]
	return t, exists
}

func DelTrack(uuid, playlist string) error {
	delete(tracks, uuid)
	pl, exists := Playlist(playlist)
	if exists {
		pl.DelTrack(uuid)
	} else {
		return ErrPlaylistNotExists
	}
	return nil
}

func Init() {
	playlists = make(map[string]*playlist.Playlist)
	tracks = make(map[string]*track.Track)
	defaultPlaylist := playlist.New()
	playlists["Default"] = defaultPlaylist
	version = "0.0"
}
