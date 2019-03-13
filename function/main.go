package main

import "fmt"

// 引入了这个包首先就会执行它的init函数
import _ "github.com/buffge/gobyexample/function/sub"

func main() {
	fmt.Println(sum0(2, 34))
	fmt.Println(sum1(2, 34))
	fmt.Println(sum1(2, 34, 0x78, 0x12))
	fmt.Println(sum1(2, 34, 34, 234, 234, 234))
	numArr := []int{1, 2, 3, 4}
	// ... 结构数组
	fmt.Println(sum1(numArr...))
	// 匿名函数
	total := sum2()
	fmt.Println(total(1))
	fmt.Println(total(2))
	fmt.Println(total(3))
	fmt.Println(total(4))
	fmt.Println(min(53, 31, 643, 124, 435, 124, 35, 32, 4, 21))
	// 在test.go 中也引入sub.go 但是因为init执行过了，所以不会再次执行
	HelloWorld()
}

/**
变量参数,返回值
*/
func sum0(a, b int) (res int) {
	res = a + b
	return
}

/**
可变参数
*/
func sum1(numArr ...int) (total int) {
	for _, num := range numArr {
		total += num
	}
	return
}

// 匿名函数

func sum2() func(b int) int {
	a := 0
	return func(b int) int {
		a = a + b
		return a
	}
}

/**
... param
*/
func min(a ...int) (m int) {
	m = int(^uint(0) >> 1)
	for _, v := range a {
		if v < m {
			m = v
		}
	}
	return
}
