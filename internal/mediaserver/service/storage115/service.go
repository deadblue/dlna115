package storage115

import "github.com/deadblue/elevengo"

type Service struct {
	ea *elevengo.Agent
}

func New() *Service {
	return &Service{
		ea: elevengo.Default(),
	}
}
