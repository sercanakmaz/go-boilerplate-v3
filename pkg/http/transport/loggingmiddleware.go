package transport

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"time"

	"go-boilerplate-v3/pkg/log"
)

type loggignMiddleware struct {
	logger             log.Logger
	ignoreRequestBody  bool
	ignoreResponseBody bool
}

func NewLoggingMiddleware(logger log.Logger) MiddlewareBase {
	return &loggignMiddleware{logger: logger}
}

func NewLoggingMiddlewareWithOptions(logger log.Logger, ignoreRequestBody bool, ignoreResponseBody bool) MiddlewareBase {
	return &loggignMiddleware{logger: logger, ignoreRequestBody: ignoreRequestBody, ignoreResponseBody: ignoreResponseBody}
}

func (self *loggignMiddleware) Request(ctx context.Context, req *http.Request) error {
	var contents = []byte{}

	if req.Body != nil && !self.ignoreRequestBody {
		contents, _ = ioutil.ReadAll(req.Body)
		req.Body = ioutil.NopCloser(bytes.NewReader(contents))
	}

	reqRelativeUrl := req.URL.Path + "?" + req.URL.RawQuery
	var withFields = log.Field{
		Message:        "Starting request",
		RequestBody:    string(contents),
		Duration:       0,
		HttpStatusCode: 102,
		Url:            reqRelativeUrl,
		HostName:       req.URL.Host,
		HttpMethod:     req.Method}

	self.logger.Request(ctx, withFields)

	return nil
}
func (self *loggignMiddleware) Response(ctx context.Context, req *http.Request, resp *http.Response) error {
	var contents = []byte{}
	var responseStatusCode int

	if resp != nil {
		responseStatusCode = resp.StatusCode

		if resp.Body != nil && !self.ignoreResponseBody {
			contents, _ = ioutil.ReadAll(resp.Body)
			resp.Body = ioutil.NopCloser(bytes.NewReader(contents))
		}
	}

	var elapsed time.Duration = ctx.Value("elapsed").(time.Duration)

	reqRelativeUrl := req.URL.Path + "?" + req.URL.RawQuery
	var withFields = log.Field{
		Message:        "Finished request",
		ResponseBody:   string(contents),
		Duration:       elapsed.Milliseconds(),
		HttpStatusCode: responseStatusCode,
		Url:            reqRelativeUrl,
		HostName:       req.URL.Host,
		HttpMethod:     req.Method}

	self.logger.Response(ctx, withFields)

	return nil
}
