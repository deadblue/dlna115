package impl

import (
	"log"

	"github.com/deadblue/dlna115/pkg/storage"
)

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
		err = s.fileFetchContent(fr, content)
	case _FetchTypeHls:
		err = s.videoFetchContent(fr, content)
	}
	if err != nil {
		content = nil
		log.Printf("Fetch content [%s] failed: %s", fr.FilePath, err)
	}
	return
}
