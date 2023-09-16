package impl

import (
	"net"
	"net/http"
	"time"

	"github.com/deadblue/dlna115/pkg/util"
	"github.com/deadblue/elevengo"
	"github.com/deadblue/elevengo/option"
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

	// Video ticket cache
	vtc *util.TTLCache[*elevengo.VideoTicket]
	// Download ticket cache
	dtc *util.TTLCache[*elevengo.DownloadTicket]
}

func New(opts *Options) (s *Service) {
	s = &Service{
		opts: opts,
		ea:   newAgent(),

		vtc: util.NewCache[*elevengo.VideoTicket](time.Minute * 30),
		dtc: util.NewCache[*elevengo.DownloadTicket](time.Minute * 30),
	}
	return
}

func newAgent() *elevengo.Agent {
	// Setup dialer
	dialer := &net.Dialer{
		Resolver: &net.Resolver{
			PreferGo: false,
		},
	}
	dialer.SetMultipathTCP(true)
	// Setup http transport
	transport := http.DefaultTransport.(*http.Transport).Clone()
	transport.DialContext = dialer.DialContext
	transport.MaxIdleConnsPerHost = 10
	// Custom http client of elevengo.Agent
	return elevengo.New(
		&option.AgentHttpOption{
			Client: &http.Client{
				Transport: transport,
			},
		},
	)
}
