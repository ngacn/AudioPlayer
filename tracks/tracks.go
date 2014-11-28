package tracks

import (
	"errors"
	"os/exec"
	"runtime"
	"strings"
)

var (
	ErrTrackNotExists = errors.New("Track not exists")

	ErrNotImplemented = errors.New("Method not implemented")
)

type Tracks map[string]*Track

// track结构中现在只有一个path， 之后还会添加信息
type Track struct {
	path     string
	playlist string
}

func (tk Track) Playlist() string {
	return tk.playlist
}

func New() Tracks {
	tks := make(map[string]*Track)
	return tks
}

func (tks Tracks) Track(uuid string) (*Track, bool) {
	tk, exists := tks[uuid]
	return tk, exists
}

func (tks Tracks) AddTrack(path, playlist string) (string, error) {
	// 使用了mac和linux自带的一个命令生成uuid
	// windows暂时没有这个＝。＝
	// 所以为了跨平台要自己实现这个生成uuid
	var uuid string
	switch os := runtime.GOOS; os {
	case "darwin", "linux":
		for {
			out, err := exec.Command("uuidgen").Output()
			if err != nil {
				return "", err
			}
			uuid = strings.TrimSpace(string(out))
			if _, exists := tks[uuid]; !exists {
				break
			}
		}
	default:
		return "", ErrNotImplemented
	}
	tk := &Track{path: path, playlist: playlist}
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
