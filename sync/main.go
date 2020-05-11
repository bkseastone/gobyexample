package main

import (
	"log"
	"sync"
)

var printOnce sync.Once

func fn1() {
	printOnce.Do(func() {
		log.Println("1. 这段话只会被输出一次")
	})
	log.Println("2. 这段话只会被输出多次")
}

func syncOnceNote() {
	log.Println("begin")
	fn1()
	fn1()
	log.Println("end")
}

func main() {
	syncOnceNote()
}
