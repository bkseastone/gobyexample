package main

import (
	"fmt"
	"time"
)

func main() {
	/**
	  初始化5个请求
	*/
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)
	//新建一个定时器 这个定时器就是NewTicker的简化版 没有close
	limiter := time.Tick(200 * time.Millisecond)

	for req := range requests {
		//速率限制
		<-limiter
		//处理这个请求
		fmt.Println("request", req, time.Now().Format("2006-01-02 15:04:05.000"))
	}
	//新建一个允许突发3个请求的限制器 其实就是加3个缓冲
	burstyLimiter := make(chan time.Time, 3)

	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}
	//每隔0.2秒给限制器发送一个请求
	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()
	//初始化5个请求
	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	//处理这5个请求,前面3个缓冲区会被立即执行掉
	// 后面的差不多每0.2秒处理一个,第一个可能会快点,取决于0.2秒发送器离现在这行代码执行的时间
	//如果这里time.Sleep(1 * time.Second) 那么 前面的4个请求会被立即执行掉,后面的每个0.2秒
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now().Format("2006-01-02 15:04:05.000"))
	}
}
