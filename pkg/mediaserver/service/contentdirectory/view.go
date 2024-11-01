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
	// Check range header
	isRange, offset, length := false, int64(0), int64(-1)
	if rh := req.Header.Get("Range"); rh != "" {
		isRange = true
		offset, length = parseRequestRange(rh)
	}
	// Fetch content
	content, err := s.ss.Fetch(filePath, offset, length)
	if err != nil {
		http.NotFound(rw, req)
		return
	}
	defer content.Body.Close()
	// Send response
	headers := rw.Header()
	headers.Set("Content-Type", content.MimeType)
	headers.Set("Content-Length", strconv.FormatInt(content.BodySize, 10))
	if cr := content.Range; isRange && cr != nil {
		headers.Set("Accept-Ranges", "bytes")
		headers.Set("Content-Range", fmt.Sprintf(
			"bytes %d-%d/%d", cr.Start, cr.End, cr.Total,
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
