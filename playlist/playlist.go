package playlist

import (
	"errors"
	"warpten/track"
)

var (
	ErrIndexOutOfRange = errors.New("Track index out of range")
)

type Playlist struct {
	tracks []*track.Track
}

func New() *Playlist {
	return new(Playlist)
}

func (pl *Playlist) Track(n int) (*track.Track, error) {
	if n >= 1 && n <= pl.Len() {
		return pl.tracks[n-1], nil
	}
	return nil, ErrIndexOutOfRange
}

func (pl *Playlist) DelTrack(uuid string) {
	for i, t := range pl.tracks {
		if t.UUID() == uuid {
			pl.tracks = append(pl.tracks[:i], pl.tracks[i+1:]...)
			break
		}
	}
}

func (pl *Playlist) Len() int {
	return len(pl.tracks)
}

func (pl *Playlist) Append(track ...*track.Track) {
	pl.tracks = append(pl.tracks, track...)
}

func (pl *Playlist) Clear() {
	pl.tracks = make([]*track.Track, 0)
}
