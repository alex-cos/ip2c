package ip2c

import (
	"fmt"

	"github.com/alex-cos/zerr"
)

var (
	ErrBadFormat = zerr.NewSC(zerr.Error, 6107, "bad returned formatted string")

	ErrUnexpected = zerr.NewSC(zerr.Error, 6108, "unexpected error")

	ErrNotFound = zerr.NewSC(zerr.Error, 6109, "IP Address was not found")

	ErrLocalhost = zerr.NewSC(zerr.Error, 6110, "can't check localhost ipaddress")
)

func ErrDoRequest(err error) error {
	msg := fmt.Sprintf("failed to execute HTTP request: %v", err)
	return zerr.NewSC(zerr.Error, 6102, msg)
}
