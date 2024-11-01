package impl

import "errors"

var (
	errInvalidPath = errors.New("invalid request path")

	errInvalidExt = errors.New("invalid request extension")

	errSegmentNotFoud = errors.New("segment not found")
)
