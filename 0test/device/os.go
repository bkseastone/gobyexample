package device

import (
	"github.com/StackExchange/wmi"
)

type operatingSystem struct {
	Name    string
	Version string
}

func GetOSInfo() ([]operatingSystem, error) {
	var operatingsystem []operatingSystem
	err := wmi.Query("Select * from Win32_OperatingSystem", &operatingsystem)
	if err != nil {
		return nil, err
	}
	return operatingsystem, nil
}
