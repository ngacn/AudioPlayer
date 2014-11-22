package player

import (
	"errors"
	"warpten/playlist"
)

var (
	ErrPlaylistExists    = errors.New("Playlist already exists")
	ErrPlaylistNotExists = errors.New("Playlist not exists")
)

var playlists map[string]*playlist.Playlist

func Playlists() *map[string]*playlist.Playlist {
	return &playlists
}

func Playlist(name string) (*playlist.Playlist, bool) {
	pl, exists := playlists[name]
	return pl, exists
}

func NewPlaylist(name string) error {
	if _, exists := playlists[name]; exists {
		return ErrPlaylistExists
	}
	playlists[name] = playlist.New(name)
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

func init() {
	playlists = make(map[string]*playlist.Playlist)
	defaultPlaylist := playlist.New("Default")
	playlists["Default"] = defaultPlaylist
}
