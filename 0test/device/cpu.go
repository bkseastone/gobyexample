package device

import (
	"github.com/StackExchange/wmi"
)

type CpuInfo struct {
	Name          string
	NumberOfCores uint32
	ThreadCount   uint32
}

func GetCPUInfo() ([]CpuInfo, error) {

	var cpuinfo []CpuInfo

	err := wmi.Query("Select * from Win32_Processor", &cpuinfo)
	if err != nil {
		return nil, err
	}
	return cpuinfo, nil
}
