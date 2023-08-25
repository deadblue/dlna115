package mediaserver

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/deadblue/dlna115/internal/mediaserver/service/connectionmanager"
	"github.com/deadblue/dlna115/internal/mediaserver/service/contentdirectory"
	"github.com/deadblue/dlna115/internal/upnp"
	"github.com/deadblue/dlna115/internal/upnp/device"
)

const (
	descUrl = "/device/desc.xml"
)

func (s *Server) initDesc(name string) {
	// Fill description
	desc := (&device.Desc{}).Init(upnp.DeviceTypeMediaServer1)
	desc.Device.UDN = s.udn
	desc.Device.FriendlyName = name
	desc.Device.Manufacturer = "deadblue"
	desc.Device.ManufacturerURL = "https://github.com/deadblue"
	desc.Device.ModelDescription = "A DLNA server implementation to stream video files from your 115 cloud storage."
	desc.Device.ModelName = "DLNA115"
	desc.Device.ModelNumber = "1.0.0"
	desc.Device.ModelURL = "https://github.com/deadblue/dlna115"
	// desc.Device.SerialNumber = ""
	// desc.Device.PresentationURL = "https://github.com/deadblue/dlna115"
	desc.Device.ServiceList.Services = []device.Service{
		{
			ServiceType: connectionmanager.ServiceType,
			ServiceId:   connectionmanager.ServiceId,
			ScpdURL:     connectionmanager.DescUrl,
			ControlURL:  connectionmanager.ControlUrl,
			EventSubURL: connectionmanager.EventUrl,
		},
		{
			ServiceType: contentdirectory.ServiceType,
			ServiceId:   contentdirectory.ServiceId,
			ScpdURL:     contentdirectory.DescUrl,
			ControlURL:  contentdirectory.ControlUrl,
			EventSubURL: contentdirectory.EventUrl,
		},
	}
	s.desc, _ = marshalXml(desc)
}

func (s *Server) handleDescXml(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "text/xml")
	rw.Header().Set("Content-Length", strconv.Itoa(len(s.desc)))
	rw.Header().Set("Server", upnp.ServerName)
	rw.WriteHeader(http.StatusOK)
	rw.Write(s.desc)
}

// ----- |upnp.Device| implementation Begin -----

func (s *Server) DeviceType() string {
	return upnp.DeviceTypeMediaServer1
}

func (s *Server) DeviceUSN() string {
	return fmt.Sprintf("%s::%s", s.udn, upnp.DeviceTypeMediaServer1)
}

func (s *Server) GetDeviceDescURL(ip string) string {
	return fmt.Sprintf(
		"http://%s:%d%s", ip, s.sp, descUrl,
	)
}

// ----- |upnp.Device| implementation End -----
