package contentdirectory

import (
	_ "embed"
	"net/http"

	"github.com/deadblue/dlna115/pkg/upnp"
	"github.com/deadblue/dlna115/pkg/util"
)

func (s *Service) HandleDescXml(rw http.ResponseWriter, req *http.Request) {
	util.SendXML(rw, upnp.ContentDirectory1Desc)
}
