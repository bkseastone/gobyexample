package main

import (
	"fmt"
	"time"
)

/**
  定时器
*/
func main() {
	timer0 := time.NewTimer(2000 * time.Millisecond)
	fmt.Println("两秒后会输出done")
	<-timer0.C
	fmt.Println("done")
	timer1 := time.NewTimer(2000 * time.Millisecond)
	fmt.Println("两秒后还会输出done")
	stopTimer := make(chan int)
	go func() {
		time.Sleep(1 * time.Second)
		stopTimer <- 0
		timer1.Stop()
	}()
	func() {
		for {
			select {
			case <-timer1.C:
				fmt.Println("done")
			case <-stopTimer:
				fmt.Println("不会输出done了,因为定时器已经被关掉了")
				return
			}
		}
	}()

	ticker0 := time.NewTicker(300 * time.Millisecond)
	go func() {
		for t := range ticker0.C {
			fmt.Println("Tick at", t.Format("2006-01-02 15:04:05.000"))
		}
	}()
	time.Sleep(3 * time.Second)
	ticker0.Stop()
	fmt.Println("定时器已关闭")
}
