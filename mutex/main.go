package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//可以创建多个锁
	var lock sync.Mutex
	var lock2 sync.Mutex
	i := 1
	var begin = make(chan int)
	go func() {
		fmt.Println("协程正在加载")
		lock.Lock()
		begin <- 0
		i++
		defer lock.Unlock()
		fmt.Println("协程还剩3秒解锁")
		time.Sleep(3 * time.Second)
	}()
	<-begin
	lock.Lock()
	fmt.Println("这段话会3秒钟后再输出,因为前一个锁要3秒后才解锁")
	fmt.Println("i的值是", i)
	lock.Unlock()
	go func() {
		fmt.Println("协程正在加载")
		lock2.Lock()
		begin <- 0
		i++
		defer lock2.Unlock()
		fmt.Println("协程还剩3秒解锁")
		time.Sleep(3 * time.Second)
	}()
	<-begin
	lock2.Lock()
	fmt.Println("这段话会3秒钟后再输出,因为前一个锁要3秒后才解锁")
	fmt.Println("i的值是", i)
	lock2.Unlock()
}
