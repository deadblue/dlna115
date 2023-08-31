package storage

type ItemType int

const (
	ItemTypeDir ItemType = iota
	ItemTypeVideo
	// TODO: Audio and Image will be supported laaaaater ...
	// ItemTypeAudio
	// ItemTypeImage
)

// Item interface for a file or dir item on storage.
type Item interface {
	Type() ItemType
}

// _BaseItem contains common fields for all types of items.
type _BaseItem struct {
	ID   string
	Name string
}

// Dir is directory item.
type Dir struct {
	_BaseItem
}

func (i *Dir) Type() ItemType {
	return ItemTypeDir
}

// Video file is video file item.
type VideoFile struct {
	_BaseItem

	// File size
	Size int64
	// Media length in seconds
	Duration float64

	// Audio channels
	AudioChannels int
	// Audio sample rate
	AudioSampleRate int

	// Video resolution
	VideoResolution string

	// URL to play
	PlayURL string
}

func (i *VideoFile) Type() ItemType {
	return ItemTypeVideo
}
