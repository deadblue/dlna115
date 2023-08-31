package upnp

import (
	"fmt"

	"github.com/deadblue/dlna115/pkg/util"
)

var (
	ServerName = fmt.Sprintf(
		"%s UPnP/1.0 DLNA115/1.0", util.OsVersion,
	)
)
