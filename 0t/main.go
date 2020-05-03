package main

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"time"
)

func main() {
	// справка
	//if len(os.Args) != 2 {
	//	fmt.Fprintf(os.Stderr, "Using: %s ip-addr\n", os.Args[0])
	//	os.Exit(1)
	//}

	target := "47.96.13.188"

	activeThreads := 0
	doneChannel := make(chan bool)

	for port := 0; port <= 65535; port++ {
		go testTCPConnection(target, port, doneChannel)
		activeThreads++
	}
	log.Println("begin", activeThreads)
	// Wait for all threads to finish
	for activeThreads > 0 {
		<-doneChannel
		activeThreads--
	}
	log.Println("end")
}

func testTCPConnection(ip string, port int, doneChannel chan bool) {
	_, err := net.DialTimeout("tcp", ip+":"+strconv.Itoa(port),
		time.Second*3)
	if err == nil {
		fmt.Printf("Port %d: Open\n", port)
	} else {
		fmt.Println("err: ", err)
	}
	doneChannel <- true
}
