package playlist

import (
	"errors"
)

var (
	ErrIndexOutOfRange = errors.New("Track index out of range")
)

type Playlists struct {
	playlists []*Playlist
}

type Playlist struct {
	name   string
	tracks []string
}

func New(name string) *Playlist {
	return &Playlist{name: name}
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
