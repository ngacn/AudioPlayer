// 所有track保存在tracks中，track属于哪个playlist只是track的一个tag，
// 所以track在各个playlist中移动几乎没有消耗
// track在daemon端只用uuid管理， track在播放列表的顺序只由client管理
package tracks

import (
	"errors"
	"warpten/utils"
)

var (
	ErrTrackNotExists = errors.New("Track not exists")
)

type Tracks map[string]*Track

// track结构中现在只有一个path， 之后还会添加信息
type Track struct {
	Path     string
	Playlist string // playlist的uuid
}

func New() Tracks {
	tks := make(map[string]*Track)
	return tks
}

func (tks Tracks) Track(uuid string) (Track, bool) {
	tk, exists := tks[uuid]
	return *tk, exists
}

func (tks Tracks) AddTrack(path, playlist string) (string, error) {
	uuid := utils.Uuidgen("track")
	tk := &Track{Path: path, Playlist: playlist}
	tks[uuid] = tk
	return uuid, nil
}

func (tks Tracks) DelTrack(uuid string) error {
	if _, exists := tks[uuid]; exists {
		delete(tks, uuid)
		return nil
	}
	return ErrTrackNotExists
}
