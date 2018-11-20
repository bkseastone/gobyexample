package main

import (
	"fmt"
	"time"
)

func main() {
	go workerPool(5)
	<-time.After(3 * time.Second)

}

func workerPool(poolNum int) {
	jobs := make(chan int, poolNum)
	res := 0
	worker := func(jobs chan int) {
		//在这里不停的接受任务并执行
		for job := range jobs {
			fmt.Println("正在执行一些操作")
			time.Sleep(300 * time.Millisecond)
			fmt.Println("操作已完成")
			res += job
		}

	}
	for i := 0; i < poolNum; i++ {
		go worker(jobs)
	}
	for i := 0; i < 10000; i++ {
		//在这里投递任务
		jobs <- i
	}
}
