package api

import (
	"context"
	"io"
	"net/http"
)

type JobCreate func(o ...func(*JobCreateRequest)) (*Response, error)

type JobCreateRequest struct {
	Body io.Reader

	header http.Header

	ctx context.Context
}
