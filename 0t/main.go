package main

import (
	"flag"
	"fmt"
	"sync"
	"time"
)

var primes = []int64{2, 3, 5, 7, 11}
var jobs = make(chan int64, 100000000)
var lock sync.Mutex

func main() {
	var max, i, j, concurrent int64

	flag.Int64Var(&max, "n", 10000, "最大数")
	flag.Int64Var(&concurrent, "c", 5, "并发数")
	flag.Parse()
	begin := time.Now()
	go func() {
		for i = 13; i < max; i += 2 {
			jobs <- i
		}
		close(jobs)
	}()
	for j = 0; j < concurrent; j++ {
		go worker()
	}
	fmt.Println(primes)
	end := time.Now()
	diff := end.Sub(begin)
	fmt.Println(diff)
}
func worker() {
	for job := range jobs {
		if isPrime(job) {
			lock.Lock()
			primes = append(primes, job)
			lock.Unlock()
		}
	}
}
func isPrime(n int64) bool {
	isPrime := true
	for _, prime := range primes {
		if n%prime == 0 {
			isPrime = false
			break
		}
	}
	return isPrime
}
