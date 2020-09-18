package main

import (
	"log"
	"net"
	"strconv"
	"sync/atomic"
	"time"
)

func main() {
	portScan("47.96.13.188")
}
func portScan(url string) {
	now := time.Now()
	// url := "www.buffge.com"
	var scanedCount int32 = 0
	openPorts := make(chan int, 66668)
	for port := 1; port < 10000; port++ {
		go func(port int) {
			mnet := new(net.Dialer)
			mnet.Timeout = 10 * time.Millisecond
			_, err := mnet.Dial("tcp", url+":"+strconv.Itoa(port))
			if err == nil {
				// log.Println(port, "端口已开放")
				openPorts <- port
			}
			atomic.AddInt32(&scanedCount, 1)
		}(port)
	}
	log.Println("正在扫描")
	for {
		currScanedCount := atomic.LoadInt32(&scanedCount)
		if currScanedCount == 65535 {
			log.Printf("扫描结束共用时%s秒\n", time.Now().Sub(now))
			break
		}

		select {
		case port := <-openPorts:
			log.Println(port, "端口已开放")
		default:
			time.Sleep(1 * time.Millisecond)
		}
	}

}
