package contentdirectory

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

func (s *Service) HandleView(rw http.ResponseWriter, req *http.Request) {
	filePath := req.URL.Path[_ViewUrlLen:]
	offset, length := parseRequestRange(req.Header.Get("Range"))

	content, err := s.ss.Fetch(filePath, offset, length)
	if err != nil {
		http.NotFound(rw, req)
		return
	}
	defer content.Body.Close()

	rw.Header().Set("Accept-Ranges", "bytes")
	rw.Header().Set("Content-Type", content.MimeType)
	rw.Header().Set("Content-Length", strconv.FormatInt(content.BodySize, 10))
	if offset != 0 || length != -1 {
		rw.Header().Set("Content-Range", fmt.Sprintf(
			"bytes %d-%d/%d",
			offset, offset+content.BodySize-1, content.FileSize,
		))
		rw.WriteHeader(http.StatusPartialContent)
	} else {
		rw.WriteHeader(http.StatusOK)
	}

	if req.Method != http.MethodHead {
		// Send content
		io.Copy(rw, content.Body)
	}
}

func parseRequestRange(rangeStr string) (offset, length int64) {
	if rangeStr == "" {
		return 0, -1
	}

	index := strings.IndexRune(rangeStr, '=')
	rangeStr = rangeStr[index+1:]
	index = strings.IndexRune(rangeStr, '-')
	startStr, endStr := rangeStr[:index], rangeStr[index+1:]

	offset, _ = strconv.ParseInt(startStr, 10, 64)
	if endStr == "" {
		length = -1
	} else {
		end, _ := strconv.ParseInt(endStr, 10, 64)
		length = end - offset + 1
	}
	return
}
