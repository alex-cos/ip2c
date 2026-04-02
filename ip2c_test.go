package ip2c_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/alex-cos/ip2c"
	"github.com/stretchr/testify/assert"
)

func TestCheck(t *testing.T) {
	t.Parallel()

	api := ip2c.NewWithTimeout(5 * time.Second)

	resp, err := api.Check("94.238.20.184")
	assert.NoError(t, err)
	if !testing.Short() {
		fmt.Printf("resp = %+v\n", resp)
	}
	assert.Equal(t, "FR", resp.CountryCode)
	assert.Equal(t, "France", resp.CountryName)

	resp, err = api.Check("111.229.215.194")
	assert.NoError(t, err)
	if !testing.Short() {
		fmt.Printf("resp = %+v\n", resp)
	}
	assert.Equal(t, "CN", resp.CountryCode)
	assert.Equal(t, "China", resp.CountryName)

	resp, err = api.Check("103.248.33.51")
	assert.NoError(t, err)
	if !testing.Short() {
		fmt.Printf("resp = %+v\n", resp)
	}
	assert.Equal(t, "IN", resp.CountryCode)
	assert.Equal(t, "India", resp.CountryName)
}

func TestCheckError(t *testing.T) {
	t.Parallel()

	api := ip2c.New()

	resp, err := api.Check("127.0.0.1")
	assert.Error(t, err)
	assert.Equal(t, (*ip2c.CheckResponseAPI)(nil), resp)
	assert.ErrorContains(t, err, "can't check localhost ipaddress")

	resp, err = api.Check("::1")
	assert.Error(t, err)
	assert.Equal(t, (*ip2c.CheckResponseAPI)(nil), resp)
	assert.ErrorContains(t, err, "can't check localhost ipaddress")

	resp, err = api.Check("abcd")
	assert.Error(t, err)
	assert.Equal(t, (*ip2c.CheckResponseAPI)(nil), resp)
	assert.ErrorContains(t, err, "invalid IP address")
}
