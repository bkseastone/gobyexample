package main

import "fmt"

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}
func main() {
	i := 1
	fmt.Println("初始化一个值:", i)
	zeroval(i)
	fmt.Println("值传递后i=:", i)
	zeroptr(&i)
	fmt.Println("指针传递后i=:", i)
	fmt.Println("&i=:", &i)
}
