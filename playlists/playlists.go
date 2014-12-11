package playlists

import (
	"errors"
	"warpten/utils"
)

var (
	ErrPlaylistNotExists = errors.New("Playlist not exists")
)

type Playlists map[string]*Playlist

type Playlist struct {
	Uuid   string   `json:"uuid"`
	Name   string   `json:"name"`
	Tracks []string `json:"tracks"` // track的uuid
}

func New() Playlists {
	pls := make(map[string]*Playlist)
	return pls
}

func (pls Playlists) Playlist(uuid string) (Playlist, bool) {
	pl, exists := pls[uuid]
	return *pl, exists
}

func (pls Playlists) AddPlaylist(name string) (Playlist, error) {
	uuid := utils.Uuidgen("playlist")
	pl := &Playlist{Uuid: uuid, Name: name, Tracks: make([]string, 0)}
	pls[uuid] = pl
	return *pl, nil
}

func (pls Playlists) DelPlaylist(uuid string) error {
	if _, exists := pls[uuid]; exists {
		pls[uuid].Tracks = make([]string, 0)
		delete(pls, uuid)
		return nil
	}
	return ErrPlaylistNotExists
}

func (pls Playlists) Clear() {
	for uuid := range pls {
		pls.DelPlaylist(uuid)
	}
}

func (pls Playlists) Len() int {
	return len(pls)
}

func (pls Playlists) AddUUIDs(pl_uuid string, tk_uuids ...string) error {
	if pl, exists := pls[pl_uuid]; exists {
		pls[pl_uuid].Tracks = append(pl.Tracks, tk_uuids...)
		return nil
	}
	return ErrPlaylistNotExists
}

func (pls Playlists) DelUUIDs(pl_uuid string, tk_uuids ...string) error {
	if pl, exists := pls[pl_uuid]; exists {
		// wtf???
		// 删除一个slice中的几个元素有没有更好的方法？
		for _, uuid := range tk_uuids {
			for i, u := range pl.Tracks {
				if uuid == u {
					pls[pl_uuid].Tracks = append(pl.Tracks[:i], pl.Tracks[i+1:]...)
					break
				}
			}
		}
		return nil
	}
	return ErrPlaylistNotExists
}
