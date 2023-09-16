package connectionmanager

import (
	"github.com/deadblue/dlna115/pkg/upnp"
)

// ----- |upup.Service| implementation Begin -----

func (s *Service) ServiceType() string {
	return upnp.ServiceTypeConnectionManager1
}

func (s *Service) ServiceId() string {
	return upnp.ServiceIdConnectionManager
}

func (s *Service) ServiceDescURL() string {
	return _DescUrl
}

func (s *Service) ServiceControlURL() string {
	return _ControlUrl
}

func (s *Service) ServiceEventURL() string {
	return _EventUrl
}

// ----- |upup.Service| implementation End -----
