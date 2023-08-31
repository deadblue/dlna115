package upnp

import _ "embed"

var (
	//go:embed desc/ConnectionManager1.xml
	ConnectionManager1Desc []byte

	//go:embed desc/ContentDirectory1.xml
	ContentDirectory1Desc []byte
)
