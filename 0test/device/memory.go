package device

import (
	"syscall"
	"unsafe"
)

var kernel = syscall.NewLazyDLL("Kernel32.dll")

type memoryStatusEx struct {
	CbSize                  uint32
	DwMemoryLoad            uint32
	UllTotalPhys            uint64 // in bytes
	UllAvailPhys            uint64
	UllTotalPageFile        uint64
	UllAvailPageFile        uint64
	UllTotalVirtual         uint64
	UllAvailVirtual         uint64
	UllAvailExtendedVirtual uint64
}

func GetMemoryInfo() (memoryStatusEx, error) {
	GlobalMemoryStatusEx := kernel.NewProc("GlobalMemoryStatusEx")
	var memInfo memoryStatusEx
	memInfo.CbSize = uint32(unsafe.Sizeof(memInfo))
	_, _, err := GlobalMemoryStatusEx.Call(uintptr(unsafe.Pointer(&memInfo)))
	return memInfo, err
}
