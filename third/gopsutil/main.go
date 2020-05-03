package main

import (
	"log"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
)

func memInfo() {
	smStat, _ := mem.SwapMemory()
	log.Println("swap mem :")
	log.Printf("总内存: %dkb\n", smStat.Total)
	log.Printf("可用内存: %dkb\n", smStat.Free)
	log.Printf("已使用: %dkb\n", smStat.Used)
	log.Printf("使用率: %.2f%%\n", smStat.UsedPercent*100)
}
func cpuInfo() {
	log.Println("cpu 信息 :")
	logicalCpuCount, _ := cpu.Counts(true)
	physicsCpuCount, _ := cpu.Counts(false)
	cpuPercentInfo, _ := cpu.Percent(0, false)
	log.Printf("逻辑cpu个数 %d\n", logicalCpuCount)
	log.Printf("物理cpu个数 %d\n", physicsCpuCount)
	log.Printf("cpu最近使用率 %.f%%\n", cpuPercentInfo[0])
}
func main() {
	memInfo()
	cpuInfo()
}
