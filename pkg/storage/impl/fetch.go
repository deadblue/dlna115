package impl

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"mime"
	"path/filepath"
	"regexp"

	"github.com/deadblue/dlna115/pkg/storage"
	"github.com/deadblue/elevengo"
)

const (
	_FetchTypeFile = "file"
	_FetchTypeHls  = "hls"

	_MimeTypeHls = "application/x-mpegURL"
)

var (
	_ReFetchPath = regexp.MustCompile(`^(\w+)-(\w+)/(\w+)\.(\w+)$`)
)

type FetchRequest struct {
	// Fetch type
	Type string

	// Original file extension
	OriginalExt string

	// Video pick code
	PickCode string

	// Request extension
	RequestExt string

	// Fetch range
	Offset, Length int64
}

func (fr *FetchRequest) Parse(path string) (err error) {
	match := _ReFetchPath.FindAllStringSubmatch(path, 1)
	if len(match) == 0 || len(match[0]) != 5 {
		return errInvalidPath
	}
	fr.Type = match[0][1]
	fr.OriginalExt = match[0][2]
	fr.PickCode = match[0][3]
	fr.RequestExt = match[0][4]
	// Check parameters
	if fr.OriginalExt != fr.RequestExt {
		return errInvalidExt
	}
	return
}

func (s *Service) Fetch(
	path string, offset int64, length int64,
) (content *storage.Content, err error) {
	// Parse fetch request
	fr := &FetchRequest{
		Offset: offset,
		Length: length,
	}
	if err = fr.Parse(path); err != nil {
		return
	}

	// Fetch content
	content = &storage.Content{}
	switch fr.Type {
	case _FetchTypeFile:
		err = s.fetchFileContent(fr, content)
	case _FetchTypeHls:
		err = s.fetchHlsContent(fr, content)
	}
	if err != nil {
		content = nil
		log.Printf("Fetch content [%s] failed: %s", fr.PickCode, err)
	}
	return
}

func (s *Service) fetchFileContent(fr *FetchRequest, content *storage.Content) (err error) {
	// Get download ticket
	ticket, ok := s.dtc.Get(fr.PickCode)
	if !ok {
		ticket = &elevengo.DownloadTicket{}
		if err = s.ea.DownloadCreateTicket(fr.PickCode, ticket); err == nil {
			s.dtc.Put(fr.PickCode, ticket)
		} else {
			return
		}
	}

	// Fetch
	content.MimeType = mime.TypeByExtension("." + fr.OriginalExt)
	content.FileSize = ticket.FileSize
	if fr.Offset == 0 && fr.Length < 0 {
		content.BodySize = content.FileSize
		content.Body, err = s.ea.Fetch(ticket.Url)
	} else {
		content.Body, err = s.ea.FetchRange(
			ticket.Url, elevengo.RangeMiddle(fr.Offset, fr.Length),
		)
		content.BodySize = content.FileSize - fr.Offset
		if fr.Length > 0 && fr.Length < content.BodySize {
			content.BodySize = fr.Length
		}
	}
	return
}

func (s *Service) fetchHlsContent(fr *FetchRequest, content *storage.Content) (err error) {
	// Get video ticket
	ticket, ok := s.vtc.Get(fr.PickCode)
	if !ok {
		ticket = &elevengo.VideoTicket{}
		if err = s.ea.VideoCreateTicket(fr.PickCode, ticket); err == nil {
			s.vtc.Put(fr.PickCode, ticket)
		} else {
			return
		}
	}

	// Read HLS data
	body, err := s.ea.Fetch(ticket.Url)
	if err != nil {
		return
	}
	defer body.Close()
	hlsData, err := io.ReadAll(body)
	if err != nil {
		return
	}

	// Make HLS content
	content.MimeType = _MimeTypeHls
	content.FileSize = int64(len(hlsData))
	content.BodySize = content.FileSize
	content.Body = io.NopCloser(bytes.NewReader(hlsData))
	return
}

func (s *Service) generatePath(file *elevengo.File) string {
	fetchType := _FetchTypeFile
	if file.IsVideo && !s.opts.DisableHLS {
		fetchType = _FetchTypeHls
	}
	fileExt := (filepath.Ext(file.Name))[1:]
	return fmt.Sprintf(
		"%s-%s/%s.%s",
		fetchType, fileExt, file.PickCode, fileExt,
	)
}
