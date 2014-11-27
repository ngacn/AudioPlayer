package playlists

import (
	"errors"
)

var (
	ErrPlaylistExists    = errors.New("Playlist already exists")
	ErrPlaylistNotExists = errors.New("Playlist not exists")

	// ErrIndexOutOfRange = errors.New("Track index out of range")
)

type Playlists map[string][]string

func New() Playlists {
	pls := make(map[string][]string)
	return pls
}

func (pls Playlists) Playlist(name string) ([]string, bool) {
	pl, exists := pls[name]
	return pl, exists
}

func (pls Playlists) AddPlaylist(name string) error {
	if _, exists := pls[name]; exists {
		return ErrPlaylistExists
	}
	pls[name] = make([]string, 0)
	return nil
}

func (pls Playlists) DelPlaylist(name string) error {
	if _, exists := pls[name]; exists {
		pls[name] = make([]string, 0)
		delete(pls, name)
		return nil
	}
	return ErrPlaylistNotExists
}

func (pls Playlists) Clear() {
	for name := range pls {
		pls.DelPlaylist(name)
	}
}

func (pls Playlists) Len() int {
	return len(pls)
}

func (pls Playlists) AddUUIDs(name string, uuids ...string) error {
	if pl, exists := pls[name]; exists {
		pl = append(pl, uuids...)
		return nil
	}
	return ErrPlaylistNotExists
}

func (pls Playlists) DelUUIDs(name string, uuids ...string) error {
	if pl, exists := pls[name]; exists {
		// wtf???
		for _, uuid := range uuids {
			for i, u := range pl {
				if uuid == u {
					pl = append(pl[:i], pl[i+1:]...)
					break
				}
			}
		}
		return nil
	}
	return ErrPlaylistNotExists
}
