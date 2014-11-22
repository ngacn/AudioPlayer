package playlist

import (
	"errors"
)

var (
	ErrIndexOutOfRange = errors.New("Track index out of range")
)

type Playlist struct {
	tracks []string
}

func New() *Playlist {
	return new(Playlist)
}

func (pl *Playlist) Track(n int) (string, error) {
	if n >= 1 && n <= pl.Len() {
		return pl.tracks[n-1], nil
	}
	return *new(string), ErrIndexOutOfRange
}

func (pl *Playlist) Len() int {
	return len(pl.tracks)
}

func (pl *Playlist) Append(track ...string) {
	pl.tracks = append(pl.tracks, track...)
}

func (pl *Playlist) Clear() {
	pl.tracks = make([]string, 0)
}
