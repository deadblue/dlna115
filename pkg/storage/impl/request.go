package impl

import "regexp"

const (
	_FetchTypeFile = "file"
	_FetchTypeHls  = "hls"
)

var (
	_ReFetchPath = regexp.MustCompile(`^(\w+)-(\w+)/(.+)\.(\w+)$`)
)

type FetchRequest struct {
	// Fetch type
	Type string

	// Original file extension
	OriginalExt string

	// Real file path
	FilePath string

	// Request file extension
	RequestExt string

	// Fetch range
	Offset, Length int64
}

func (fr *FetchRequest) Parse(path string) (err error) {
	matches := _ReFetchPath.FindStringSubmatch(path)
	if len(matches) == 0 {
		return errInvalidPath
	}
	fr.Type = matches[1]
	fr.OriginalExt = matches[2]
	fr.FilePath = matches[3]
	fr.RequestExt = matches[4]
	return
}
