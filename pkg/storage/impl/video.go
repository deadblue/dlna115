package impl

import (
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/deadblue/dlna115/pkg/compress"
	"github.com/deadblue/dlna115/pkg/storage"
	"github.com/deadblue/dlna115/pkg/util"
	"github.com/deadblue/elevengo"
	"github.com/grafov/m3u8"
)

type VideoSegment struct {
	Url      string
	Duration float64
}

type VideoMetadata struct {
	// HLS target duration
	TargetDuration float64
	// HLS segments
	Segments []*VideoSegment
	// Headers to download HLS segment
	Headers map[string]string
}

func (s *Service) videoFetchContent(fr *FetchRequest, content *storage.Content) (err error) {
	parts := strings.SplitN(fr.FilePath, "/", 3)
	if len(parts) == 1 {
		if fr.OriginalExt != fr.RequestExt {
			return errInvalidExt
		}
		return s.videoCreateManifest(parts[0], content)
	} else {
		if fr.RequestExt != "ts" {
			return errInvalidExt
		}
		return s.videoFetchSegment(parts[2], content)
	}
}

func (s *Service) videoCreateManifest(pickcode string, content *storage.Content) (err error) {
	vm, ok := s.vmc.Get(pickcode)
	if !ok {
		vm = &VideoMetadata{
			Headers: make(map[string]string),
		}
		if err = s.videoFetchMetadata(pickcode, vm); err != nil {
			return
		}
		// Cache for 6 hours
		s.vmc.Put(pickcode, vm, 6*time.Hour)
	}
	// Construct m3u8 content
	sb := &strings.Builder{}
	sb.WriteString("#EXTM3U\n")
	sb.WriteString("#EXT-X-VERSION:3\n")
	sb.WriteString("#EXT-X-ALLOW-CACHE:YES\n")
	sb.WriteString(fmt.Sprintf("#EXT-X-TARGETDURATION:%d\n", int(vm.TargetDuration)))
	sb.WriteString("#EXT-X-MEDIA-SEQUENCE:0\n")
	for index, segment := range vm.Segments {
		sb.WriteString(fmt.Sprintf(
			"#EXTINF:%0.6f,\n%s/%d/%s.ts\n",
			segment.Duration, pickcode, index, compress.Encode(segment.Url),
		))
	}
	sb.WriteString("#EXT-X-ENDLIST\n")
	// Fill storage.Content
	body := strings.NewReader(sb.String())
	content.Body = io.NopCloser(body)
	content.BodySize = body.Size()
	content.MimeType = util.MimeTypeM3U8
	return
}

func (s *Service) videoFetchMetadata(pickcode string, vm *VideoMetadata) (err error) {
	// Get video ticket
	ticket := &elevengo.VideoTicket{}
	if err = s.ea.VideoCreateTicket(pickcode, ticket); err != nil {
		return
	}
	for name, value := range ticket.Headers {
		vm.Headers[name] = value
	}
	// Fetch master playlist
	var pl m3u8.Playlist
	if pl, err = s.videoFetchHlsPlaylist(ticket.Url); err != nil {
		return
	}
	// Select highest media playlist
	masterPl := pl.(*m3u8.MasterPlaylist)
	bestVariant := (*m3u8.Variant)(nil)
	for _, variant := range masterPl.Variants {
		if bestVariant == nil || bestVariant.Bandwidth < variant.Bandwidth {
			bestVariant = variant
		}
	}
	// Fetch media playlist
	if pl, err = s.videoFetchHlsPlaylist(bestVariant.URI); err == nil {
		mediaPl := pl.(*m3u8.MediaPlaylist)
		vm.TargetDuration = mediaPl.TargetDuration
		for _, segment := range mediaPl.Segments {
			if segment == nil {
				break
			}
			vm.Segments = append(vm.Segments, &VideoSegment{
				Duration: segment.Duration,
				// Get absolute URL
				Url: util.GetAbsoluteUrl(bestVariant.URI, segment.URI),
			})
		}
	}
	return
}

func (s *Service) videoFetchHlsPlaylist(url string) (pl m3u8.Playlist, err error) {
	body, err := s.ea.Fetch(url)
	if err != nil {
		return
	}
	defer body.Close()
	pl, _, err = m3u8.DecodeFrom(body, true)
	return
}

func (s *Service) videoFetchSegment(name string, content *storage.Content) (err error) {
	segmentUrl, err := compress.Decode(name)
	if err != nil {
		return
	}
	body, err := s.ea.Fetch(segmentUrl)
	if err == nil {
		// TODO: Support range?
		content.Body = body
		content.BodySize = body.Size()
		content.MimeType = util.MimeTypeM2TS
	}
	return
}
