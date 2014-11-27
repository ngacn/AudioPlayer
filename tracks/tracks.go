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

type Track struct {
	path string
}

func New() Tracks {
	tks := make(map[string]*Track)
	return tks
}

func (tks Tracks) Track(uuid string) (*Track, bool) {
	tk, exists := tks[uuid]
	return tk, exists
}

func (tks Tracks) AddTrack(path string) (string, error) {
	var uuid string
	switch os := runtime.GOOS; os {
	case "darwin":
	case "linux":
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
	tk := &Track{path: path}
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
