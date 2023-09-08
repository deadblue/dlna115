package impl

import (
	"time"

	"github.com/deadblue/dlna115/pkg/util"
	"github.com/deadblue/elevengo"
)

type Folder struct {
	// Folder display name
	Name string
	// Folder type
	Type string
	// SourceId should be dirId or labelId
	SourceId string
}

type Service struct {
	// 115 options
	opts *Options
	// 115 agent
	ea *elevengo.Agent
	// Top folders
	tfs []*Folder
	// Download ticket cache
	dtc *util.TTLCache[*elevengo.DownloadTicket]
}

func New(opts *Options) (s *Service) {
	s = &Service{
		opts: opts,
		ea:   elevengo.Default(),
		dtc:  util.NewCache[*elevengo.DownloadTicket](time.Minute * 30),
	}
	return
}
