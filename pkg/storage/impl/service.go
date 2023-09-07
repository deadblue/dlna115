package impl

import (
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
}

func New(opts *Options) (s *Service) {
	s = &Service{
		opts: opts,
		ea:   elevengo.Default(),
	}
	return
}
