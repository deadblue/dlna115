package contentdirectory

import (
	"github.com/deadblue/dlna115/pkg/upnp"
)

// ----- |upup.Service| implementation Begin -----

func (s *Service) ServiceType() string {
	return upnp.ServiceTypeContentDirectory1
}

func (s *Service) ServiceId() string {
	return upnp.ServiceIdContentDirectory
}

func (s *Service) ServiceDescURL() string {
	return descUrl
}

func (s *Service) ServiceControlURL() string {
	return controlUrl
}

func (s *Service) ServiceEventURL() string {
	return eventUrl
}

// ----- |upup.Service| implementation End -----
