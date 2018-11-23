package device

import (
	"github.com/StackExchange/wmi"
)

type gpuInfo struct {
	Name string
}

func GetGPUInfo() ([]gpuInfo, error) {

	var myGpuinfo []gpuInfo
	err := wmi.Query("Select * from Win32_VideoController", &myGpuinfo)
	if err != nil {
		return []gpuInfo{}, err
	}
	return myGpuinfo, nil
}
