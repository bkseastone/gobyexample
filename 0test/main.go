package main

import (
	"fmt"
	"github.com/buffge/gobyexample/0test/device"
)

func main() {
	cpuInfo, _ := device.GetCPUInfo()
	//fmt.Printf("%#v\n", cpuInfo)
	fmt.Printf("cpu个数 %v\n", len(cpuInfo))
	fmt.Printf("cpu名称 %v\n", cpuInfo[0].Name)
	fmt.Printf("cpu核心数 %v\n", cpuInfo[0].NumberOfCores)
	fmt.Printf("cpu线程数 %v\n", cpuInfo[0].ThreadCount)
	osIfno, _ := device.GetOSInfo()
	//fmt.Printf("%#v\n", osIfno)
	fmt.Printf("系统个数 %v\n", len(osIfno))
	fmt.Printf("系统名称 %v\n", osIfno[0].Name)
	fmt.Printf("系统版本 %v\n", osIfno[0].Version)
	memIfno, _ := device.GetMemoryInfo()
	//fmt.Printf("%#v\n", osIfno)
	fmt.Printf("系统位数 %v\n", memIfno.CbSize)
	fmt.Printf("系统已使用内存 %.2f\n", float64(memIfno.DwMemoryLoad/1024/1024/1024))
	fmt.Printf("有效扩展内存U %.2f\n", float64(memIfno.UllAvailExtendedVirtual/1024/1024/1024))
	fmt.Printf("闲置页面文件大小 %.2f\n", float64(memIfno.UllAvailPageFile/1024/1024/1024))
	fmt.Printf("闲置物理内存 %.2f\n", float64(memIfno.UllAvailPhys/1024/1024/1024))
	fmt.Printf("闲置虚拟内存 %.2f\n", float64(memIfno.UllAvailVirtual/1024/1024/1024))
	fmt.Printf("总内存 %.2f\n", float64(memIfno.UllTotalPageFile)/1024/1024/1024)
	fmt.Printf("总物理内存 %.2f\n", float64(memIfno.UllTotalPhys/1024/1024/1024))
	fmt.Printf("总虚拟内存 %.2f\n", float64(memIfno.UllTotalVirtual/1024/1024/1024))
}
