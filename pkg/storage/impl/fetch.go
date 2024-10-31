package impl

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/deadblue/dlna115/pkg/storage"
	"github.com/deadblue/elevengo"
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
