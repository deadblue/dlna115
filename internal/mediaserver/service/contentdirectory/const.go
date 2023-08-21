package contentdirectory

const (
	ServiceType = "urn:schemas-upnp-org:service:ContentDirectory:1"
	ServiceId   = "urn:upnp-org:serviceId:ContentDirectory"

	DescUrl    = "/ContentDirectory/scpd.xml"
	ControlUrl = "/ContentDirectory/control"
	EventUrl   = "/ContentDirectory/event"

	ActionBrowse = "Browse"
	ActionSearch = "Search"
)
