package main

import (
	"context"
	"fmt"
	"log"
	"runtime"
	"strings"
	"sync"

	"golang.org/x/sync/semaphore"
)

var printOnce sync.Once
var a sync.Cond
var sbPool = &sync.Pool{
	New: func() interface{} {
		return interface{}(&strings.Builder{})
	},
}

func fn1() {
	printOnce.Do(func() {
		log.Println("1. 这段话只会被输出一次")
	})
	log.Println("2. 这段话只会被输出多次")
}

func syncOnceNote() {
	log.Println("begin")
	fn1()
	log.Println("end")
}
func testSyncOnce() {
	syncOnceNote()
	syncOnceNote()
}
func testSbPool() {
	sb := sbPool.Get().(*strings.Builder)
	sb.WriteString("hello")
	log.Println(sb.String()) // Hello
	sbPool.Put(sb)
	sb2 := sbPool.Get().(*strings.Builder)
	log.Println(sb2.String()) // Hello
	sb.Reset()
	sbPool.Put(sb)
	sb3 := sbPool.Get().(*strings.Builder)
	log.Println(sb3.String()) // 空字符串
}
func testErrGroup() {}
func testSemaphore() {
	ctx := context.Background()
	var (
		maxWorkers = runtime.GOMAXPROCS(0)
		// 一共maxWorkers个信号量
		sem = semaphore.NewWeighted(int64(maxWorkers))
		out = make([]int, 32)
	)
	// 并行maxWorker计算
	for i := range out {
		// 每个协程占一个信号量(逻辑cpu)
		if err := sem.Acquire(ctx, 1); err != nil {
			log.Printf("Failed to acquire semaphore: %v", err)
			break
		}
		// 协程工作完后释一个信号量(逻辑cpu)
		go func(i int) {
			defer sem.Release(1)
			// doSomething
			out[i] = i + 1
		}(i)
	}
	// 确保上面的工作全部完成（即全部释放）
	if err := sem.Acquire(ctx, int64(maxWorkers)); err != nil {
		log.Printf("Failed to acquire semaphore: %v", err)
	}

	fmt.Println(out)
}
func main() {
	testSemaphore()
}
