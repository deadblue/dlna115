package storage115

import (
	"fmt"

	"github.com/deadblue/dlna115/internal/mediaserver/service/storageservice"
	"github.com/deadblue/elevengo"
)

func (s *Service) Browse(parentId string) (items []storageservice.Item) {
	items = make([]storageservice.Item, 0)
	it, err := s.ea.FileIterate(parentId)
	if err != nil {
		return
	}
	// Travel files
	file := elevengo.File{}
	for ; err == nil; err = it.Next() {
		it.Get(&file)
		if file.IsDirectory {
			item := &storageservice.Dir{}
			item.ID = file.FileId
			item.Name = file.Name
			// Append to items
			items = append(items, item)
		} else if file.IsVideo {
			item := &storageservice.VideoFile{}
			item.ID = file.FileId
			item.Name = file.Name
			item.Size = file.Size
			item.Duration = file.MediaDuration
			// Make play URL
			item.PlayURL = fmt.Sprintf("%s%s.m3u8", VideoURL, file.PickCode)
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
			// Append to items
			items = append(items, item)
		}
	}
	return
}
