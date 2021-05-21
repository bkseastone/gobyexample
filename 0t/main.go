package main

import (
	"log"
)

func f1() error {
	defer func() {
		if err := recover(); err != nil {
			log.Println("recover")
		}
	}()
	panic(123)
}

func main() {
	var a map[int]int
	log.Println(a[1])
}
