package upnp

import (
	"fmt"

	"github.com/deadblue/dlna115/internal/util"
)

var (
	ServerName = fmt.Sprintf(
		"%s UPnP/1.0 DLNA115/1.0", util.OsVersion,
	)
)
