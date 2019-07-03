package main

import (
	"fmt"
	"log"
	"time"
)

func main() {

	c0 := make(chan int)
	c1 := make(chan int)
	timeout := time.After(time.Second * 3)
	go func() {
		time.AfterFunc(time.Second*1, func() {
			c0 <- 0
		})
		time.AfterFunc(time.Second*2, func() {
			c1 <- 0
		})
	}()
	go func() {
		select {
		case <-c0:
			println("c0先收到值")
		case <-c1:
			println("c1先收到值")
		case <-timeout:
			println("已超时")
		}
	}()

	//缓冲通道
	c2 := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			c2 <- i
			time.Sleep(time.Millisecond * 100)
		}
		close(c2)
	}()
	log.Println("begin")
	//在这里会一直卡住,直到通道被关闭
	for v := range c2 {
		fmt.Println("接收到chan值", v)
	}
	log.Println("end")
	//通道通知
	workerComplete := make(chan byte)
	workerTimeout := time.After(time.Millisecond * 1000)
	go func() {
		fmt.Println("执行一些耗时的操作,比如查询数据库,请求网络api等等")
		time.Sleep(time.Millisecond * 800)
		workerComplete <- 0
	}()
	select {
	case <-workerComplete:
		fmt.Println("耗时操作完成了")
	case <-workerTimeout:
		fmt.Println("两秒还没执行完,操作超时了")
	}

	//单向通道
	ping := make(chan byte)
	pong := make(chan byte)
	// 这里一秒后发送消息到ping中
	go func(ping chan<- byte) {
		time.Sleep(time.Second * 1)
		ping <- 22
	}(ping)
	go func(ping <-chan byte, pong chan<- byte) {
		// 一秒后ping输出,pong接收到新消息
		v := <-ping
		pong <- v
	}(ping, pong)
	select {
	case <-pong:
		log.Println("已接收到ping信息,即将发出pong信息")
	}
	messages := make(chan string)
	msg := "hi"
	/**
	这段代码就让messages 在另一个协程输出
	go func() {
		<-messages
	}()
	time.Sleep(10 * time.Millisecond)
	*/
	select {
	// 因为messages 没有缓冲区 所以 msg 不能发进来 如果开个协程用 <-messages 就可以了
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}
}
