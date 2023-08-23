package ssdp

import (
	"fmt"
	"io"
)

var (
	crlf = []byte("\r\n")
)

type Request struct {
	Method  string
	Headers map[string]string
}

func (r *Request) SetHeader(name, value string) {
	if r.Headers == nil {
		r.Headers = make(map[string]string)
	}
	r.Headers[name] = value
}

func (r *Request) WriteTo(w io.Writer) (n64 int64, err error) {
	var n int
	startLine := fmt.Sprintf("%s * HTTP/1.1\r\n", r.Method)
	if n, err = w.Write([]byte(startLine)); err != nil {
		return
	}
	n64 += int64(n)
	for name, value := range r.Headers {
		headerLine := fmt.Sprintf("%s: %s\r\n", name, value)
		if n, err = w.Write([]byte(headerLine)); err != nil {
			return
		}
		n64 += int64(n)
	}
	if n, err = w.Write(crlf); err != nil {
		return
	}
	n64 += int64(n)
	return
}

type Response struct {
	StatusCode int
	StatusText string
	Headers    map[string]string
}

func (r *Response) SetHeader(name, value string) {
	if r.Headers == nil {
		r.Headers = make(map[string]string)
	}
	r.Headers[name] = value
}
