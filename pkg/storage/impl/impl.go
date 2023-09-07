package impl

import (
	"fmt"
	"strings"

	"github.com/deadblue/dlna115/pkg/storage"
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
	switch parts[0] {
	case FolderTypeDir:
		return s.browseDir(parts[1])
	case FolderTypeStar:
		return s.browseStar()
	case FolderTypeLabel:
		return s.browseLabel(parts[1])
	}
	return emptyItems
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

func (s *Service) browseDir(dirId string) []storage.Item {
	it, err := s.ea.FileIterate(dirId)
	if err != nil {
		return emptyItems
	} else {
		return s.createItemList(it)
	}
}

func (s *Service) browseStar() []storage.Item {
	it, err := s.ea.FileWithStar()
	if err != nil {
		return emptyItems
	} else {
		return s.createItemList(it)
	}
}

func (s *Service) browseLabel(labelId string) []storage.Item {
	it, err := s.ea.FileWithLabel(labelId)
	if err != nil {
		return emptyItems
	} else {
		return s.createItemList(it)
	}
}

func (s *Service) createItemList(it elevengo.Iterator[elevengo.File]) []storage.Item {
	items := make([]storage.Item, 0)
	var err error
	for ; err == nil; err = it.Next() {
		file := &elevengo.File{}
		if it.Get(file) != nil {
			continue
		}
		if file.IsDirectory {
			items = append(items, createDir(file))
		} else if file.IsVideo {
			items = append(items, createVideoFile(file, s.opts.DisableHLS))
		}
	}
	return items
}

func createDir(file *elevengo.File) (item *storage.Dir) {
	item = &storage.Dir{}
	item.Name = file.Name
	item.ID = fmt.Sprintf("%s@%s", FolderTypeDir, file.FileId)
	return
}

func createVideoFile(file *elevengo.File, disableHLS bool) (item *storage.VideoFile) {
	item = &storage.VideoFile{}
	item.ID = file.FileId
	item.Name = file.Name
	item.Size = file.Size
	item.Duration = file.MediaDuration
	// Make play URL
	playType := PlayTypeStream
	if disableHLS {
		playType = PlayTypeFile
	}
	item.PlayURL = fmt.Sprintf(
		"%s%s/%s", VideoURL, playType, file.PickCode,
	)
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
