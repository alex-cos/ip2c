package ip2c

import "context"

type IP2C interface {
	Check(ipAddress string) (*CheckResponseAPI, error)
	CheckWithContext(ctx context.Context, ipAddress string) (*CheckResponseAPI, error)
}
