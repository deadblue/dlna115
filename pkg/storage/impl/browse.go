package impl

import (
	"fmt"
	"log"
	"path/filepath"
	"strings"

	"github.com/deadblue/dlna115/pkg/storage"
	"github.com/deadblue/dlna115/pkg/util"
	"github.com/deadblue/elevengo"
)

var (
	emptyItems = make([]storage.Item, 0)
)

func (s *Service) Browse(parentId string) (items []storage.Item) {
	if parentId == "0" {
		return s.browseRoot()
	}
	parts := strings.SplitN(parentId, "@", 2)

	var err error
	var it elevengo.Iterator[elevengo.File]
	switch parts[0] {
	case FolderTypeDir:
		it, err = s.ea.FileIterate(parts[1])
	case FolderTypeStar:
		it, err = s.ea.FileWithStar()
	case FolderTypeLabel:
		it, err = s.ea.FileWithLabel(parts[1])
	}
	if err != nil {
		log.Printf("Get file list failed: %s", err)
		return emptyItems
	} else {
		return s.createItemList(it)
	}
}

func (s *Service) browseRoot() (items []storage.Item) {
	items = make([]storage.Item, len(s.tfs))
	for i, tf := range s.tfs {
		item := &storage.Dir{}
		item.Name = tf.Name
		item.ID = fmt.Sprintf("%s@%s", tf.Type, tf.SourceId)
		items[i] = item
	}
	return
}

func (s *Service) createItemList(it elevengo.Iterator[elevengo.File]) []storage.Item {
	items := make([]storage.Item, 0)
	for _, file := range it.Items() {
		if file.IsDirectory {
			items = append(items, s.createDir(file))
		} else if file.IsVideo && file.VideoDefinition > 0 {
			items = append(items, s.createVideoFile(file))
		} else if !file.IsVideo && file.MediaDuration > 0 {
			items = append(items, s.createAudioFile(file))
		} else if util.IsImageFile(file.Name) {
			items = append(items, s.createImageFile(file))
		}
	}
	return items
}

func (s *Service) createDir(file *elevengo.File) (item *storage.Dir) {
	item = &storage.Dir{}
	item.ID = fmt.Sprintf("%s@%s", FolderTypeDir, file.FileId)
	item.Name = file.Name
	return
}

func (s *Service) createVideoFile(file *elevengo.File) (item *storage.VideoFile) {
	item = &storage.VideoFile{}
	item.ID = file.FileId
	item.Name = file.Name
	item.Size = file.Size
	item.MimeType = util.GetMimeType(file.Name)
	item.URLPath = s.generateFilePath(file)
	item.Duration = file.MediaDuration
	// GUESS resoltion from video definition
	switch file.VideoDefinition {
	case elevengo.VideoDefinitionSD:
		item.VideoResolution = "640x480"
	case elevengo.VideoDefinitionHD:
		item.VideoResolution = "1280x720"
	case elevengo.VideoDefinitionFHD, elevengo.VideoDefinition1080P:
		item.VideoResolution = "1920x1080"
	case elevengo.VideoDefinition4K:
		item.VideoResolution = "3840x2160"
	default:
		// Fallback
		item.VideoResolution = "640x480"
	}
	// Dummy values which we can not get form 115
	item.AudioChannels = 2
	item.AudioSampleRate = 44100
	return
}

func (s *Service) createAudioFile(file *elevengo.File) (item *storage.AudioFile) {
	item = &storage.AudioFile{}
	item.ID = file.FileId
	item.Name = file.Name
	item.Size = file.Size
	item.MimeType = util.GetMimeType(file.Name)
	item.URLPath = s.generateFilePath(file)
	item.Duration = file.MediaDuration
	// Dummy values which we can not get form 115
	item.AudioChannels = 2
	item.AudioSampleRate = 44100
	return
}

func (s *Service) createImageFile(file *elevengo.File) (item *storage.ImageFile) {
	item = &storage.ImageFile{}
	item.ID = file.FileId
	item.Name = file.Name
	item.Size = file.Size
	item.MimeType = util.GetMimeType(file.Name)
	item.URLPath = s.generateFilePath(file)
	return
}

func (s *Service) generateFilePath(file *elevengo.File) string {
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
