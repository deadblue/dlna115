package forward

import "github.com/deadblue/elevengo"

type Service struct {
	ea *elevengo.Agent
}

func New(ea *elevengo.Agent) *Service {
	return &Service{ea: ea}
}
