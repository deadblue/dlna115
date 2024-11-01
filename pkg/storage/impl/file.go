package impl

import (
	"net/url"
	"strconv"
	"time"

	"github.com/deadblue/dlna115/pkg/storage"
	"github.com/deadblue/dlna115/pkg/util"
	"github.com/deadblue/elevengo"
)

func getFileUrlExpiration(downloadUrl string) time.Duration {
	u, _ := url.Parse(downloadUrl)
	t := u.Query().Get("t")
	expireAt, _ := strconv.ParseInt(t, 10, 64)
	expireTime := time.Unix(expireAt, 0)
	return time.Until(expireTime)
}

func (s *Service) fileFetchContent(fr *FetchRequest, content *storage.Content) (err error) {
	if fr.OriginalExt != fr.RequestExt {
		return errInvalidExt
	}
	// Get download ticket from cache or upstream
	pickcode := fr.FilePath
	ticket, ok := s.dtc.Get(pickcode)
	if !ok {
		ticket = &elevengo.DownloadTicket{}
		if err = s.ea.DownloadCreateTicket(pickcode, ticket); err != nil {
			return
		}
		s.dtc.Put(pickcode, ticket, getFileUrlExpiration(ticket.Url))
	}

	// Fetch stream from storage
	if fr.Offset == 0 && fr.Length < 0 {
		content.Body, err = s.ea.Fetch(ticket.Url)
	} else {
		content.Body, err = s.ea.FetchRange(
			ticket.Url, elevengo.RangeMiddle(fr.Offset, fr.Length),
		)
	}
	if err != nil {
		return
	}
	// Fill content metadata
	maxLength := ticket.FileSize - fr.Offset
	if fr.Length < 0 || fr.Length > maxLength {
		content.BodySize = maxLength
	} else {
		content.BodySize = fr.Length
	}
	content.MimeType = util.GetMimeTypeForExt(fr.OriginalExt)
	// Always fill range metadata
	content.Range = &storage.ContentRange{
		Start: fr.Offset,
		End:   fr.Offset + content.BodySize - 1,
		Total: ticket.FileSize,
	}
	return
}
