package ip2c

import (
	"context"
	"fmt"
	"io"
	"mime"
	"net/http"
	"strings"
	"time"
)

const (
	// APIURL is the official ip2c API Endpoint.
	APIURL = "http://ip2c.org"
)

// ----------------------------------------------------------------------------
// Structures
// ----------------------------------------------------------------------------

// IP2CAPI represents an ip2c IP2CAPI Client connection.
type IP2CAPI struct {
	client  *http.Client
	timeout time.Duration
}

func New() IP2C {
	return NewWithClient(http.DefaultClient)
}

func NewWithTimeout(timeout time.Duration) IP2C {
	return NewWithClient(&http.Client{
		Transport:     nil,
		CheckRedirect: nil,
		Jar:           nil,
		Timeout:       timeout,
	})
}

func NewWithClient(httpClient *http.Client) IP2C {
	return &IP2CAPI{
		client:  httpClient,
		timeout: httpClient.Timeout,
	}
}

func NewWithClientTimeout(
	httpClient *http.Client,
	timeout time.Duration,
) IP2C {
	return &IP2CAPI{
		client:  httpClient,
		timeout: timeout,
	}
}

func (api *IP2CAPI) Check(ipAddress string) (*CheckResponseAPI, error) {
	if isLocalHost(ipAddress) {
		return nil, ErrLocalhost
	}

	reqURL := fmt.Sprintf("%s/%s", APIURL, ipAddress)
	resp, err := api.doRequest(http.MethodGet, reqURL)
	if err != nil {
		return nil, err
	}
	fields := strings.Split(resp, ";")
	if len(fields) < 4 {
		return nil, ErrBadFormat
	}
	switch fields[0] {
	case "0":
		return nil, ErrUnexpected
	case "1":
		return &CheckResponseAPI{
			CountryCode: strings.ToUpper(fields[1]),
			CountryName: fields[3],
		}, nil
	case "2":
		return nil, ErrNotFound
	default:
		return nil, ErrUnexpected
	}
}

// Unexported functions

func (api *IP2CAPI) doRequest(method, reqURL string) (string, error) {
	ctx, cancel := api.getContext()
	if cancel != nil {
		defer cancel()
	}
	req, err := http.NewRequestWithContext(ctx, method, reqURL, nil)
	if err != nil {
		return "", ErrNewRequest(err)
	}
	resp, err := api.client.Do(req)
	if err != nil {
		return "", ErrDoRequest(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", ErrReadBody(method, reqURL, err)
	}

	mimeType, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
	if err != nil {
		return "", ErrParseContentType(method, reqURL, err)
	}
	if mimeType == "text/html" {
		if len(body) == 0 {
			return "", ErrEmptyBody(method, reqURL)
		}
		return string(body), nil
	}
	return "", ErrUnsupportedMimeType(method, reqURL, mimeType)
}

func (api *IP2CAPI) getContext() (context.Context, context.CancelFunc) {
	var cancel context.CancelFunc

	ctx := context.Background()
	if api.timeout > 0 {
		ctx, cancel = context.WithTimeout(context.Background(), api.timeout)
	}

	return ctx, cancel
}

func isLocalHost(ip string) bool {
	return (ip == "127.0.0.1") || (ip == "::1")
}
