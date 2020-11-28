package main

import (
	"log"
	"strings"
	"sync"
)

var printOnce sync.Once

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
func main() {
	testSbPool()
}
