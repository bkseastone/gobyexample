package main

import (
	"fmt"
	"net"
	"strconv"
	"sync/atomic"
	"time"
)

func portScan() {
	now := time.Now()
	url := "www.buffge.com"
	var scanedCount int32 = 0
	openPorts := make(chan int, 66668)
	for port := 1; port < 65536; port++ {
		go func(port int) {
			mnet := new(net.Dialer)
			mnet.Timeout = 1000 * time.Millisecond
			_, err := mnet.Dial("tcp", url+":"+strconv.Itoa(port))
			if err == nil {
				//fmt.Println(port, "端口已开放")
				openPorts <- port
			}
			atomic.AddInt32(&scanedCount, 1)
		}(port)
	}
	fmt.Println("正在扫描")
	for {
		currScanedCount := atomic.LoadInt32(&scanedCount)
		if currScanedCount == 65535 {
			fmt.Printf("扫描结束共用时%s秒\n", time.Now().Sub(now))
			break
		}

		select {
		case port := <-openPorts:
			fmt.Println(port, "端口已开放")
		default:
			time.Sleep(1 * time.Millisecond)
		}
	}

}
