package playlists

import "testing"

func TestPlaylist(t *testing.T) {
	const (
		track1 string = "/home/foo/Music/Rurutia/1.mp3"
		track2 string = "/home/foo/Music/Rurutia/2.mp3"
	)

	pl := New()
	pl.Append(track1, track2)

	if l := pl.Len(); l != 2 {
		t.Errorf("Len() = %v, want %v", l, 2)
	}

	if track, err := pl.Track(0); err != ErrIndexOutOfRange {
		t.Errorf("Track(0) = %v, %v, want %v", track, err, ErrIndexOutOfRange)
	}

	if track, err := pl.Track(1); err != nil || track != track1 {
		t.Errorf("Track(1) = %v, %v, want %v", track, err, track1)
	}

	if track, err := pl.Track(2); err != nil || track != track2 {
		t.Errorf("Track(2) = %v, %v, want %v", track, err, track2)
	}

	if track, err := pl.Track(3); err != ErrIndexOutOfRange {
		t.Errorf("Track(3) = %v, %v, want %v", track, err, ErrIndexOutOfRange)
	}

	pl.Clear()

	if track, err := pl.Track(1); err != ErrIndexOutOfRange {
		t.Errorf("Track(1) = %v, %v, want %v", track, err, ErrIndexOutOfRange)
	}
}
