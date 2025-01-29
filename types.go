package ip2c

import "fmt"

type CheckResponseAPI struct {
	CountryCode string `json:"country_code"`
	CountryName string `json:"country_name"`
}

func (item *CheckResponseAPI) String() string {
	return fmt.Sprintf("%+v", *item)
}
