package main

import (
	"log"
	"syscall"
	"unsafe"
)

const (
	dingdingProcessName = "DingTalk.exe"
	// dingdingProcessName = "PVFTool.exe"
	maxPath = 260
)

func main() {
	var pid int
	targetName := dingdingProcessName
	processList, err := processes()
	if err != nil {
		log.Fatalln("获取进程列表失败")
	}
	for _, process := range processList {
		if process.Executable() == targetName {
			pid = process.Pid()
			break
		}
	}
	if pid == 0 {
		log.Fatalln("未找到进程: ", targetName)
	}
	hProcess, err := syscall.OpenProcess(syscall.PROCESS_QUERY_INFORMATION, false, uint32(pid))
	if err != nil {
		log.Fatalln("获取进程handle失败", err)
	}
	buf := [maxPath]uint16{}
	size := maxPath
	ret, _, lastErr := QueryFullProcessImageNameW.Call(uintptr(hProcess), 0, uintptr(unsafe.Pointer(&buf)), uintptr(unsafe.Pointer(&size)))
	if ret != 1 {
		log.Println("获取进程全路径失败:", lastErr)
	} else {
		log.Println("进程全路径: ", syscall.UTF16ToString(buf[:]))
	}

}
