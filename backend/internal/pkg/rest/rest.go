package rest

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"

	"go.elastic.co/apm/v2"

	"github.com/go-resty/resty/v2"
	"github.com/pebruwantoro/monorepo_project/backend/internal/pkg/logger"
)

type RestClient interface {
	Get(ctx context.Context, path string, header http.Header) (body []byte, statusCode int, err error)
	Post(ctx context.Context, path string, header http.Header, requestBody []byte) (body []byte, statusCode int, err error)
	Put(ctx context.Context, path string, header http.Header, requestBody []byte) (body []byte, statusCode int, err error)
	Patch(ctx context.Context, path string, header http.Header, requestBody []byte) (body []byte, statusCode int, err error)
	Delete(ctx context.Context, path string, header http.Header, requestBody []byte) (body []byte, statusCode int, err error)
}

type client struct {
	options    Options
	httpClient *resty.Client
}

func New(opt Options) RestClient {
	defaultClient := resty.New().
		SetTimeout(opt.Timeout)

	if opt.SkipTLS {
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}

		defaultClient = defaultClient.SetTransport(tr)
	}

	return &client{
		httpClient: defaultClient,
		options:    opt,
	}
}

func (c *client) Get(ctx context.Context, path string, header http.Header) (body []byte, statusCode int, err error) {
	url := c.options.Address + path
	body, statusCode, err = c.call(ctx, http.MethodGet, url, header, nil)
	return
}

func (c *client) Post(ctx context.Context, path string, header http.Header, requestBody []byte) (body []byte, statusCode int, err error) {
	url := c.options.Address + path
	body, statusCode, err = c.call(ctx, http.MethodPost, url, header, requestBody)
	return
}

func (c *client) Put(ctx context.Context, path string, header http.Header, requestBody []byte) (body []byte, statusCode int, err error) {
	url := c.options.Address + path
	body, statusCode, err = c.call(ctx, http.MethodPut, url, header, requestBody)
	return
}

func (c *client) Patch(ctx context.Context, path string, header http.Header, requestBody []byte) (body []byte, statusCode int, err error) {
	url := c.options.Address + path
	body, statusCode, err = c.call(ctx, http.MethodPatch, url, header, requestBody)
	return
}

func (c *client) Delete(ctx context.Context, path string, header http.Header, requestBody []byte) (body []byte, statusCode int, err error) {
	url := c.options.Address + path
	body, statusCode, err = c.call(ctx, http.MethodDelete, url, header, requestBody)
	return
}

func (c *client) call(ctx context.Context, method, path string, requestHeader http.Header, requestBody []byte,
) (body []byte, status int, err error) {
	span, ctx := apm.StartSpan(ctx, fmt.Sprintf("%s %s", method, path), "custom")

	defer func() {
		if err != nil {
			if logger.Log != nil {
				logger.Log.Error(ctx, err.Error())
			}
		}

		span.End()
	}()

	// append x-client-id header
	if c.options.ClientID != "" {
		requestHeader.Add("X-Client-Id", c.options.ClientID)
	}

	requestHeader.Del("Accept-Encoding")

	// repopulate header, because resty cannot read http.Header
	headers := make(map[string][]string)
	for key, value := range requestHeader {
		headers[key] = value
	}

	resp, err := c.httpClient.R().
		SetHeaderMultiValues(headers).
		SetBody(requestBody).
		Execute(method, path)

	if logger.Log != nil {
		logger.Log.Info(ctx, fmt.Sprintf("[Request Header] %s", method), path, requestHeader)

		if !logger.IsSkipLog(resp.Request.Header.Get("Content-Type")) {
			logger.Log.Info(ctx, fmt.Sprintf("[Request] %s", method), path, string(requestBody))
		} else {
			logger.Log.Info(ctx, "Request Not Log Because Unsupported Content-Type")
		}
	}

	body = resp.Body()
	status = resp.StatusCode()

	if logger.Log != nil {
		if !logger.IsSkipLog(resp.Header().Get("Content-Type")) {
			logger.Log.Info(ctx, fmt.Sprintf("[Response] %s", method), path, string(body))
		} else {
			logger.Log.Info(ctx, "Response Not Log Because Unsupported Content-Type")
		}
	}

	return
}
