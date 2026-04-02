package ip2c

import (
	"errors"
	"fmt"
)

var (
	ErrBadFormat = errors.New("bad returned formatted string")

	ErrUnexpected = errors.New("unexpected error")

	ErrNotFound = errors.New("IP Address was not found")

	ErrLocalhost = errors.New("can't check localhost ipaddress")

	ErrInvalidIP = errors.New("invalid IP address")
)

func ErrDoRequest(err error) error {
	return fmt.Errorf("failed to execute HTTP request: %w", err)
}
