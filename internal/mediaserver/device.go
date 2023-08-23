package mediaserver

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/deadblue/dlna115/internal/mediaserver/proto"
	"github.com/deadblue/dlna115/internal/mediaserver/service/connectionmanager"
	"github.com/deadblue/dlna115/internal/mediaserver/service/contentdirectory"
	"github.com/deadblue/dlna115/internal/upnp"
)

const (
	deviceDescUrl = "/device/desc.xml"
)

func makeDeviceDesc(uuid string) []byte {
	// Fill description
	desc := (&proto.Description{}).Init()
	desc.Device.UDN = fmt.Sprintf("uuid:%s", uuid)
	desc.Device.FriendlyName = "DLNA115"
	desc.Device.SerialNumber = uuid
	desc.Device.ModelName = "DLNA115"
	desc.Device.ModelURL = "https://github.com/deadblue/dlna115"
	desc.Device.ModelDescription = "A DLNA server implementation to stream video files from your 115 cloud storage."
	desc.Device.Manufacturer = "deadblue"
	desc.Device.ManufacturerURL = "https://github.com/deadblue"
	desc.Device.ServiceList.Services = []proto.Service{
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
	data, _ := marshalXml(desc)
	return data
}

func (s *Server) handleDescDeviceXml(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "text/xml")
	rw.Header().Set("Content-Length", strconv.Itoa(len(s.descXml)))
	rw.Header().Set("Server", upnp.ServerTag)
	rw.WriteHeader(http.StatusOK)
	rw.Write(s.descXml)
}
