package contentdirectory

import (
	_ "embed"
	"net/http"

	"github.com/deadblue/dlna115/internal/util"
)

//go:embed assets/ContentDirectory1.xml
var descXml []byte

func (s *Service) HandleDescXml(rw http.ResponseWriter, req *http.Request) {
	util.SendXML(rw, descXml)
}
