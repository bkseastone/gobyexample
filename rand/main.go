package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("因为没有设置种子所以,每次都是同样的81和87 ")
	fmt.Print("[0,100)", rand.Intn(100), ",")
	fmt.Println(rand.Intn(100))
	fmt.Println(rand.Float64())
	fmt.Print("[0.0,1,0)", (rand.Float64()*5)+5, ",")
	fmt.Print((rand.Float64() * 5) + 5)
	fmt.Println()
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	fmt.Print("设置种子后 ", r1.Intn(100), ",")
	fmt.Print(r1.Intn(100))
	fmt.Println()
	fmt.Println("也可以全局执行rand.Seed(int64) 设置种子")
	s2 := rand.NewSource(42)
	r2 := rand.New(s2)
	fmt.Print("如果是相同的种子那么,生产也是一样的", r2.Intn(100), ",")
	fmt.Print(r2.Intn(100))
	fmt.Println()
	s3 := rand.NewSource(42)
	r3 := rand.New(s3)
	fmt.Print(r3.Intn(100), ",")
	fmt.Print(r3.Intn(100))
}
