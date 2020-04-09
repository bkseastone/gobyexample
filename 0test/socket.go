package main

import (
	. "fmt"
	"strconv"
	"strings"
	"syscall"
	"time"
)

func MAKEWORD(low, high uint8) uint32 {
	var ret uint16 = uint16(high)<<8 + uint16(low)
	return uint32(ret)
}

func inet_addr(ipaddr string) [4]byte {
	var (
		ips = strings.Split(ipaddr, ".")
		ip  [4]uint64
		ret [4]byte
	)
	for i := 0; i < 4; i++ {
		ip[i], _ = strconv.ParseUint(ips[i], 10, 8)
	}
	for i := 0; i < 4; i++ {
		ret[i] = byte(ip[i])
	}
	return ret
}

func main2() {
	var (
		//sock syscall.Handle
		//addr    syscall.SockaddrInet4
		wsadata syscall.WSAData
		err     error
	)
	if err = syscall.WSAStartup(MAKEWORD(2, 2), &wsadata); err != nil {
		Println("Startup error")
		return
	}
	defer syscall.WSACleanup()

	//addr.Addr = inet_addr("47.96.13.188")
	success := make(chan int, 1000)
	for port := 1; port < 65536; port++ {
		var sock syscall.Handle

		if sock, err = syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, syscall.IPPROTO_IP); err != nil {
			Println("Socket create error")
			return
		}
		//defer syscall.Closesocket(sock)
		go func(port int) {
			var addr syscall.SockaddrInet4
			addr.Addr = inet_addr("47.96.13.188")
			addr.Port = port
			//println(time.Now().Format("2006-01-02 15:04:05.000000"))
			timeout := time.After(100 * time.Millisecond)
			openPort := make(chan int)
			notOpen := make(chan int)
			go func() {
				if err = syscall.Connect(sock, &addr); err != nil {
					//println(time.Now().Format("2006-01-02 15:04:05.000000"))
					//Println("Connect error", err)
					notOpen <- 0
					return
				} else {
					openPort <- port
					return
				}
			}()
			for {
				select {
				case <-timeout:
					syscall.Closesocket(sock)
					return
				case currPort := <-openPort:
					success <- currPort
					syscall.Closesocket(sock)
					return
				case <-notOpen:
					syscall.Closesocket(sock)
					return
				default:
					time.Sleep(10 * time.Millisecond)
				}
			}

		}(port)
	}
	timout := time.After(10 * time.Second)
	openCount := 0

	for {
		select {
		case port := <-success:
			openCount++
			println("已开放端口", port)
		case <-timout:
			println("一共开放了", openCount, "个端口")
			return
		default:
			time.Sleep(10 * time.Millisecond)
		}
	}
	//var (
	//    data       syscall.WSABuf
	//    sendstr    string = "hello"
	//    SendButes  uint32
	//    overlapped syscall.Overlapped
	//)
	//data.Len = uint32(len(sendstr))
	//data.Buf, _ = syscall.BytePtrFromString(sendstr)
	//err = syscall.WSASend(sock, &data, 1, &SendButes, 0, &overlapped, nil)
	//if err != nil {
	//    Println("Send error")
	//} else {
	//    Println("Send success")
	//}

}
