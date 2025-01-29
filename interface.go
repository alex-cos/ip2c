package ip2c

type IP2C interface {
	// Check endpoint accepts a single IP address (v4 or v6).
	Check(ipAddress string) (*CheckResponseAPI, error)
}
