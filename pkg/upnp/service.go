package upnp

const (
	ServiceTypeConnectionManager1 = "urn:schemas-upnp-org:service:ConnectionManager:1"
	ServiceTypeContentDirectory1  = "urn:schemas-upnp-org:service:ContentDirectory:1"

	ServiceIdConnectionManager = "urn:upnp-org:serviceId:ConnectionManager"
	ServiceIdContentDirectory  = "urn:upnp-org:serviceId:ContentDirectory"
)

// Service interface should be impled by an UPnP service.
type Service interface {
	// DeviceType returns UPnP service type that the service impled.
	ServiceType() string
	// ServiceId returns UPnP service id of the service.
	ServiceId() string
	// ServiceDescURL returns desc URL of the service.
	ServiceDescURL() string
	// ServiceDescURL returns control URL of the service.
	ServiceControlURL() string
	// ServiceDescURL returns event URL of the service.
	ServiceEventURL() string
}
