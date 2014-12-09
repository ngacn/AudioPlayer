package tracks

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

var (
	ErrTrackNotExists = errors.New("Track not exists")

	ErrGenerateFailed = errors.New("Fail to generate UUID")
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
func Uuidgen(r *rand.Rand) string { //基于随机数的UUID生成器，Linux默认的
	return fmt.Sprintf("%x%x-%x-%x-%x-%x%x%x",
		r.Int31(), r.Int31(),
		r.Int31(),
		(r.Int31()&0x0fff)|0x4000, //Generates a 32-bit Hex number of the form 4xxx (4 indicates the UUID version)
		r.Int31()%0x3fff+0x8000,   //range [0x8000, 0xbfff]
		r.Int31(), r.Int31(), r.Int31())
}
func (tks Tracks) AddTrack(path, playlist string) (string, error) {

	var uuid string
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	uuid = Uuidgen(r)
	if len(uuid) == 0 {
		return "", ErrGenerateFailed
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
