package main

import (
	"context"
	"fmt"
	"time"
)

func testCancel() {
	// cancel
	ctx, cancel := context.WithCancel(context.Background())
	go work(ctx, "work1")
	go work(ctx, "work2")
	go work(ctx, "work3")
	time.Sleep(time.Second * 3)
	cancel()
	time.Sleep(time.Second * 1)
}
func testValue() {
	// with value
	ctx1, valueCancel := context.WithCancel(context.Background())
	valueCtx := context.WithValue(ctx1, "name", "buffge")
	valueCtx = context.WithValue(valueCtx, "age", 24)
	go workWithValue(valueCtx, "value work", "name")
	time.Sleep(time.Second * 3)
	valueCancel()
}
func testTimeout() {
	// timeout
	ctx2, timeCancel := context.WithTimeout(context.Background(), time.Second*3)
	go work(ctx2, "time cancel")
	time.Sleep(time.Second * 5)
	timeCancel()
}
func testDeadline() {
	// deadline
	ctx3, deadlineCancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second*3))
	go work(ctx3, "deadline cancel")
	time.Sleep(time.Second * 5)
	deadlineCancel()

	time.Sleep(time.Second * 3)
}
func main() {
	testTimeout()
	//testCancel()
	// ch 一次只能通知一个 ctx.Done() 可以通知全部
	//ch := make(chan bool)
	//go testMultiChan(ch, "1")
	//go testMultiChan(ch, "2")
	//go testMultiChan(ch, "3")
	//go func(ch chan<- bool) {
	//	<-time.After(3 * time.Second)
	//	ch <- true
	//}(ch)
	time.Sleep(10 * time.Second)

}
func testMultiChan(ch <-chan bool, name string) {
	for {
		select {
		case <-ch:
			fmt.Printf("%s接收到ch,正在关闭\n", name)
			return
		default:
			<-time.After(1 * time.Second)
		}
	}
}
func workWithValue(ctx context.Context, name string, key string) {
	for {
		select {
		case <-ctx.Done():
			println(name, " get message to quit")
			return
		default:
			fmt.Printf("name = %s\n", ctx.Value("name"))
			fmt.Printf("age = %d\n", ctx.Value("age"))
			println(name, " is running", time.Now().String())
			time.Sleep(time.Second)
		}
	}
}

func work(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			println(name, " get message to quit")
			return
		default:
			println(name, " is running", time.Now().String())
			time.Sleep(time.Second)
		}
	}
}
