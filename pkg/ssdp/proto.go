package ssdp

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"net/http"
)

var (
	crlf = []byte("\r\n")
)

// SSDP request
type Request struct {
	// Request method
	Method string
	// Request headers
	headers map[string]string
}

func (req *Request) SetHeader(name, value string) {
	if req.headers == nil {
		req.headers = make(map[string]string)
	}
	req.headers[name] = value
}

func (req *Request) GetHeader(name string) string {
	if req.headers == nil {
		return ""
	}
	return req.headers[name]
}

func (req *Request) WriteTo(w io.Writer) (n int64, err error) {
	bw := bufio.NewWriterSize(w, 1500)
	startLine := fmt.Sprintf("%s * HTTP/1.1\r\n", req.Method)
	if _, err = bw.WriteString(startLine); err != nil {
		return
	}
	for name, value := range req.headers {
		headerLine := fmt.Sprintf("%s: %s\r\n", name, value)
		if _, err = bw.WriteString(headerLine); err != nil {
			return
		}
	}
	if _, err = bw.Write(crlf); err != nil {
		return
	}
	n = int64(bw.Buffered())
	err = bw.Flush()
	return
}

func (req *Request) Unmarshal(b []byte) (err error) {
	lines := bytes.Split(b, crlf)
	if len(lines) == 0 {
		return ErrMalformedRequest
	}
	// Parse start line
	if pos := bytes.IndexRune(lines[0], ' '); pos > 0 {
		req.Method = string(lines[0][:pos])
	} else {
		return ErrMalformedRequest
	}
	// Parse headers
	for _, line := range lines[1:] {
		pos := bytes.IndexRune(line, ':')
		if pos < 0 {
			continue
		}
		name := bytes.TrimSpace(line[:pos])
		value := bytes.TrimSpace(line[pos+1:])
		req.SetHeader(string(name), string(value))
	}
	return
}

// SSDP response
type Response struct {
	// Response status
	StatusCode int
	// Response headers
	headers map[string]string
}

func (r *Response) SetHeader(name, value string) {
	if r.headers == nil {
		r.headers = map[string]string{}
	}
	r.headers[name] = value
}

func (r *Response) WriteTo(w io.Writer) (n int64, err error) {
	bw := bufio.NewWriterSize(w, 1500)
	// Write start line
	startLine := fmt.Sprintf(
		"HTTP/1.1 %d %s\r\n",
		r.StatusCode, http.StatusText(r.StatusCode),
	)
	if _, err = bw.WriteString(startLine); err != nil {
		return
	}
	// Write headers
	for name, value := range r.headers {
		headerLine := fmt.Sprintf("%s: %s\r\n", name, value)
		if _, err = bw.WriteString(headerLine); err != nil {
			return
		}
	}
	if _, err = bw.Write(crlf); err != nil {
		return
	}
	n = int64(bw.Buffered())
	err = bw.Flush()
	return
}
