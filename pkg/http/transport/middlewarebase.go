package transport

import (
	"context"
	"net/http"
)

type MiddlewareBase interface {
	Request(ctx context.Context, req *http.Request) error
	Response(ctx context.Context, req *http.Request, resp *http.Response) error
}
