package ip2c

import (
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/alex-cos/restc"
)

const (
	// APIURL is the official ip2c API Endpoint.
	APIURL = "https://ip2c.org"
)

// ----------------------------------------------------------------------------
// Structures
// ----------------------------------------------------------------------------

// IP2CAPI represents an ip2c IP2CAPI Client connection.
type IP2CAPI struct {
	client *restc.Client
}

func New() IP2C {
	return NewWithClient(http.DefaultClient)
}

func NewWithClient(httpClient *http.Client) IP2C {
	return NewWithClientTimeout(httpClient, restc.DefaultTimeout)
}

func NewWithTimeout(timeout time.Duration) IP2C {
	return NewWithClientTimeout(http.DefaultClient, timeout)
}

func NewWithClientTimeout(
	httpClient *http.Client,
	timeout time.Duration,
) IP2C {
	return &IP2CAPI{
		client: restc.NewWithClientTimeout(APIURL, httpClient, timeout),
	}
}

func (api *IP2CAPI) Check(ipAddress string) (*CheckResponseAPI, error) {
	if net.ParseIP(ipAddress) == nil {
		return nil, ErrInvalidIP
	}

	if isLocalHost(ipAddress) {
		return nil, ErrLocalhost
	}

	req := restc.Get(ipAddress)
	resp, err := api.client.Execute(req)
	if err != nil {
		return nil, ErrDoRequest(err)
	}

	fields := strings.Split(string(resp.Bytes()), ";")
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

func isLocalHost(ip string) bool {
	parsed := net.ParseIP(ip)
	return parsed != nil && parsed.IsLoopback()
}
