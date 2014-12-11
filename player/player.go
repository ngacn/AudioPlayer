// 主要的播放器功能都会在这实现
package player

import (
	"warpten/playlists"
	"warpten/tracks"
)

// 播放器版本
var version string

// 所有播放列表, 以播放列表uuid为key， Playlist结构为值
var pls playlists.Playlists

// 所有track， 以uuid为key, Track结构为值
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

func Playlist(uuid string) (playlists.Playlist, bool) {
	return pls.Playlist(uuid)
}

func AddPlaylist(name string) (playlists.Playlist, error) {
	return pls.AddPlaylist(name)
}

func DelPlaylist(uuid string) error {
	pl, exists := pls.Playlist(uuid)
	if !exists {
		return playlists.ErrPlaylistNotExists
	}

	// 删除所有播放列表中uuid对应的track
	for _, tk_uuid := range pl.Tracks {
		if err := tks.DelTrack(tk_uuid); err != nil {
			return err
		}
	}

	// 删除播放列表中的uuid
	if err := pls.DelPlaylist(uuid); err != nil {
		return err
	}
	return nil
}

func Track(uuid string) (tracks.Track, bool) {
	tk, exists := tks.Track(uuid)
	return tk, exists
}

func AddTrack(path, playlist string) (tracks.Track, error) {
	// 创建新track, 并获得新track的uuid
	tk, err := tks.AddTrack(path, playlist)
	if err != nil {
		return tk, err
	}

	// 将uuid添加到对应的播放列表
	_, exists := pls.Playlist(playlist)
	if exists {
		return tk, pls.AddUUIDs(playlist, tk.Uuid)
	}
	return tk, playlists.ErrPlaylistNotExists
}

func DelTrack(uuid string) error {
	tk, exists := tks.Track(uuid)
	if !exists {
		return tracks.ErrTrackNotExists
	}
	playlist := tk.Playlist
	// 删除uuid对应的track
	if err := tks.DelTrack(uuid); err != nil {
		return err
	}

	// 删除播放列表中的uuid
	_, exists = pls.Playlist(playlist)
	if exists {
		return pls.DelUUIDs(playlist, uuid)
	}
	return playlists.ErrPlaylistNotExists
}

func New() {
	version = "0.0"
	// 初始化播放列表， 至少有一个叫Default的列表
	pls = playlists.New()
	pls.AddPlaylist("Default")
	// 初始化track列表
	tks = tracks.New()
}
