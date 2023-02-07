package transport

import (
	"context"
	"net/http"
	"time"
)

// This type implements the http.RoundTripper interface
type MiddlewareTransport struct {
	Middlewares []MiddlewareBase
	Transport   http.RoundTripper
}

func (self MiddlewareTransport) RoundTrip(req *http.Request) (res *http.Response, e error) {
	var (
		err  error
		resp *http.Response
	)
	ctx := req.Context()

	for _, middleware := range self.Middlewares {
		if err = middleware.Request(ctx, req); err != nil {
			return nil, err
		}
	}

	timer := time.Now()

	resp, err = self.Transport.RoundTrip(req)

	elapsed := time.Since(timer)

	ctx = context.WithValue(ctx, "elapsed", elapsed)

	for _, middleware := range self.Middlewares {
		if middlewareErr := middleware.Response(ctx, req, resp); middlewareErr != nil {
			return nil, middlewareErr
		}
	}

	return resp, err
}

func NewMiddlewareTransport(middlewares []MiddlewareBase, transport http.RoundTripper) *MiddlewareTransport {
	return &MiddlewareTransport{Middlewares: middlewares, Transport: transport}
}
