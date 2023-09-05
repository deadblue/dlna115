package upnp

import (
	"fmt"

	"github.com/deadblue/dlna115/pkg/util"
	"github.com/deadblue/dlna115/pkg/version"
)

var (
	ServerName = fmt.Sprintf(
		"%s UPnP/1.0 %s",
		util.OsVersion(), version.Short(),
	)
)
