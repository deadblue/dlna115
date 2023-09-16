package storage

import "io"

type ItemType int

const (
	ItemTypeDir ItemType = iota
	ItemTypeImage
	ItemTypeAudio
	ItemTypeVideo
)

// Item interface should be implemented by all Dir and File items.
type Item interface {
	Type() ItemType
}

// _BaseItem contains common fields for all types of item.
type _BaseItem struct {
	ID   string
	Name string
}

// Dir represents a folder.
type Dir struct {
	_BaseItem
}

func (i *Dir) Type() ItemType {
	return ItemTypeDir
}

type _BaseFile struct {
	_BaseItem
	// File size
	Size int64
	// Access URLPath
	URLPath string
	// MIME type
	MimeType string
}

// ImageFile represnts an image file
type ImageFile struct {
	_BaseFile
}

func (i *ImageFile) Type() ItemType {
	return ItemTypeImage
}

// AudioFile represents an audio file
type AudioFile struct {
	_BaseFile

	// Media length in seconds
	Duration float64
	// Audio channels
	AudioChannels int
	// Audio sample rate
	AudioSampleRate int
}

func (i *AudioFile) Type() ItemType {
	return ItemTypeAudio
}

// VideoFile represents a video file
type VideoFile struct {
	AudioFile

	// Video resolution
	VideoResolution string
}

func (i *VideoFile) Type() ItemType {
	return ItemTypeVideo
}

type Content struct {
	// Content body
	Body io.ReadCloser
	// Body size
	BodySize int64
	// File size
	FileSize int64
	// MIME type
	MimeType string
}
