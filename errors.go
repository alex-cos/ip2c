package ip2c

import (
	"fmt"

	"github.com/alex-cos/zerr"
)

var (
	ErrNewRequest = func(err error) error {
		msg := fmt.Sprintf("failed to build HTTP request: %v", err)
		return zerr.NewSC(zerr.Error, 6101, msg)
	}

	ErrDoRequest = func(err error) error {
		msg := fmt.Sprintf("failed to execute HTTP request: %v", err)
		return zerr.NewSC(zerr.Error, 6102, msg)
	}

	ErrReadBody = func(method, url string, err error) error {
		msg := fmt.Sprintf("failed to Read HTTP response body '%s:%s': %v",
			method, url, err)
		return zerr.NewSC(zerr.Error, 6103, msg)
	}

	ErrParseContentType = func(method, url string, err error) error {
		msg := fmt.Sprintf("failed to parse HTTP content type '%s:%s': %v",
			method, url, err)
		return zerr.NewSC(zerr.Error, 6104, msg)
	}

	ErrEmptyBody = func(method, url string) error {
		msg := fmt.Sprintf("unexpected empty body response '%s:%s'",
			method, url)
		return zerr.NewSC(zerr.Error, 6105, msg)
	}

	ErrUnsupportedMimeType = func(method, url, mimeType string) error {
		msg := fmt.Sprintf("unsupported mime type '%s:%s': %s",
			method, url, mimeType)
		return zerr.NewSC(zerr.Error, 6106, msg)
	}

	ErrBadFormat = zerr.NewSC(zerr.Error, 6107, "bad returned formatted string")

	ErrUnexpected = zerr.NewSC(zerr.Error, 6108, "unexpected error")

	ErrNotFound = zerr.NewSC(zerr.Error, 6109, "IP Address was not found")

	ErrLocalhost = zerr.NewSC(zerr.Error, 6110, "can't check localhost ipaddress")
)
