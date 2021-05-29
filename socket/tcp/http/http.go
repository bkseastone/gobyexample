package main

import (
	"bytes"
	"errors"
	"io"
	"log"
	"net"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

var packetFree = sync.Pool{
	New: func() interface{} {
		return make([]byte, 1<<12)
	},
}

func newPacket() []byte {
	return packetFree.Get().([]byte)
}

func newConn(addr string) net.Conn {
	var conn net.Conn
	var err error
	for {
		conn, err = net.Dial("tcp", addr)
		if err != nil {
			log.Println("连接失败", err)
			time.Sleep(time.Millisecond)
			continue
		}
		break
	}
	// tcpConn, _ := conn.(*net.TCPConn)
	// file, _ := tcpConn.File()
	// err = syscall.SetsockoptInt(syscall.Handle(int(file.Fd())), syscall.SOL_SOCKET, syscall.TCP_NODELAY, 1)
	// _ = file.Close()
	return conn
}

var addr string
var reqBts []byte

func spider(contentLenC chan int, failC chan struct{}, doneC chan struct{}) {
	conn := newConn(addr)
	var packet []byte
	var contentLen int
	var n int
	var err error
	go func() {
		for {
			select {
			case <-doneC:
				return
			default:
			}
			contentLen = 0
			// log.Println("in req", string(reqBts))
			_, err = conn.Write(reqBts)
			if err != nil {
				failC <- struct{}{}
				log.Println("req failed", err)
				_ = conn.Close()
				conn = newConn(addr)
			}
			// log.Println("req success")
		}
	}()
	packet = newPacket()
	for {
		select {
		case <-doneC:
			return
		default:
		}
		// log.Println("in read")
		n, err = conn.Read(packet)
		if err != nil {
			if errors.Is(err, io.EOF) {
				_ = conn.Close()
				conn = newConn(addr)
				break
			}
			log.Println("read err:", err)
			break
		}
		contentLen += n
		if contentLen > 0 {
			contentLenC <- contentLen
		} else {
			failC <- struct{}{}
		}
	}
	packetFree.Put(packet)
	_ = conn.Close()
}
func main() {
	go http.ListenAndServe(":6060", nil)
	addr = "localnor.com:8080"
	// addr = "127.0.0.1:8080"
	duration := time.Minute * 10
	reqBuf := bytes.Buffer{}
	for i := 0; i < 1000; i++ {
		reqBuf.WriteString("GET / HTTP/1.1\r\n")
		reqBuf.WriteString("Host:a.b\r\n")
		reqBuf.WriteString("\r\n")
	}
	reqBts = reqBuf.Bytes()
	concurrent := 2000
	contentLen := 0
	successCount := 0
	failedCount := 0
	contentLenC := make(chan int, concurrent*5)
	failC := make(chan struct{}, concurrent)
	doneC := make(chan struct{})
	for i := 0; i < concurrent; i++ {
		go spider(contentLenC, failC, doneC)
	}
	outputContentLenTicker := time.NewTicker(time.Second * 5)
	begin := time.Now()
	go func() {
		var duration time.Duration
		for t := range outputContentLenTicker.C {
			if contentLen == 0 {
				continue
			}
			duration = t.Sub(begin)
			lenMb := float64(contentLen / 1024 / 1024)
			durationSec := duration.Seconds()
			log.Printf(" ticker 共用时%s\n 成功%d次\n 失败%d次\n 总字节数%.2fM\n"+
				" 每秒字节数:%.1fM\n 每秒请求数: %.1f\n",
				duration, successCount, failedCount,
				lenMb,
				lenMb/durationSec,
				float64(successCount)/durationSec,
			)
		}
	}()
	quit := time.After(duration)
	signQuit := make(chan os.Signal, 1)
	signal.Notify(signQuit, syscall.SIGINT, syscall.SIGTERM)
	for {
		select {
		case t := <-contentLenC:
			successCount++
			contentLen += t
		case <-failC:
			failedCount++
		case <-signQuit:
			close(doneC)
		case <-quit:
			close(doneC)
		case <-doneC:
			goto end
		default:
			time.Sleep(10 * time.Millisecond)
		}
	}
end:
	duration = time.Since(begin)
	lenMb := float64(contentLen / 1024 / 1024)
	durationSec := duration.Seconds()
	log.Printf("访问%s\n 共用时%v\n 成功%d次\n 失败%d次\n 总字节数%.1fM\n"+
		" 每秒字节数:%.1fM\n 每秒请求数: %.1f\n",
		addr, duration, successCount, failedCount, lenMb,
		lenMb/durationSec,
		float64(successCount)/durationSec,
	)
	log.Println("end pressure")
	if outputContentLenTicker != nil {
		outputContentLenTicker.Stop()
	}

}
