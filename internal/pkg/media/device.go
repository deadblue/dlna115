package media

import (
	"fmt"

	"github.com/deadblue/dlna115/internal/pkg/media/proto"
)

const (
	connectionManagerServiceType = "urn:schemas-upnp-org:service:ConnectionManager:1"
	connectionManagerServiceId   = "urn:upnp-org:serviceId:ConnectionManager"

	contentDirectoryServiceType = "urn:schemas-upnp-org:service:ContentDirectory:1"
	contentDirectoryServiceId   = "urn:upnp-org:serviceId:ContentDirectory"
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
			ServiceType: connectionManagerServiceType,
			ServiceId:   connectionManagerServiceId,
			ScpdURL:     connectionManagerDescUrl,
			ControlURL:  connectionManagerControlUrl,
			EventSubURL: connectionManagerEventUrl,
		},
		{
			ServiceType: contentDirectoryServiceType,
			ServiceId:   contentDirectoryServiceId,
			ScpdURL:     contentDirectoryDescUrl,
			ControlURL:  contentDirectoryControlUrl,
			EventSubURL: contentDirectoryEventUrl,
		},
	}
	data, _ := marshalXml(desc)
	return data
}
