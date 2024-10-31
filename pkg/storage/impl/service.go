package impl

import (
	"net"
	"net/http"

	"github.com/deadblue/dlna115/pkg/cache"
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

	// Video metadata cache
	vmc *cache.TTLCache[*VideoMetadata]
	// Download ticket cache
	dtc *cache.TTLCache[*elevengo.DownloadTicket]
}

func New(opts *Options) (s *Service) {
	s = &Service{
		opts: opts,
		ea:   newAgent(),

		vmc: cache.New[*VideoMetadata](),
		dtc: cache.New[*elevengo.DownloadTicket](),
	}
	return
}

func newAgent() *elevengo.Agent {
	// Setup dialer
	dialer := &net.Dialer{
		Resolver: &net.Resolver{
			PreferGo: true,
		},
	}
	dialer.SetMultipathTCP(true)
	// Setup http transport
	transport := http.DefaultTransport.(*http.Transport).Clone()
	transport.DialContext = dialer.DialContext
	transport.MaxIdleConnsPerHost = 10
	// Custom http client of elevengo.Agent
	return elevengo.New(
		option.Agent().WithHttpClient(&http.Client{
			Transport: transport,
		}),
	)
}
