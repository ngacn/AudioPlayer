package track

type Track struct {
	uuid string
	path string
}

func New(uuid, path string) *Track {
	return &Track{uuid, path}
}

func (t *Track) UUID() string {
	return t.uuid
}
