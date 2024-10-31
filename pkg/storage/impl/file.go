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
	// Get download ticket
	pickcode := fr.FilePath
	ticket, ok := s.dtc.Get(pickcode)
	if !ok {
		ticket = &elevengo.DownloadTicket{}
		if err = s.ea.DownloadCreateTicket(pickcode, ticket); err == nil {
			s.dtc.Put(pickcode, ticket, getFileUrlExpiration(ticket.Url))
		} else {
			return
		}
	}

	// Fetch
	content.MimeType = util.GetMimeTypeForExt(fr.OriginalExt)
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
