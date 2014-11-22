package player

import (
	"errors"
	"warpten/playlist"
)

var (
	ErrPlaylistExists    = errors.New("Playlist already exists")
	ErrPlaylistNotExists = errors.New("Playlist not exists")
)

var version string
var playlists map[string]*playlist.Playlist

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

func Init() {
	playlists = make(map[string]*playlist.Playlist)
	defaultPlaylist := playlist.New()
	playlists["Default"] = defaultPlaylist
	version = "0.0"
}
