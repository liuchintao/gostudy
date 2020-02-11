package api

import (
	"io"
	"net/http"
)

type Response struct {
	StatusCode int
	Header     http.Header
	Body       io.ReadCloser
}

func (r *Response)
