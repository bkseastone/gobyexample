package main

import "fmt"

/**
 * go程序中有 字符串 整数 浮点数 布尔值
	bool

	string

	int  int8  int16  int32  int64
	uint uint8 uint16 uint32 uint64 uintptr

	byte // uint8 的别名

	rune // int32 的别名
		 // 代表一个Unicode码

	float32 float64

	complex64 complex128
*/
func main() {
	fmt.Println("this is " + "string")
	fmt.Println("this is int", 234)
	fmt.Println("this is float", 2.2423)
	fmt.Println("this is bool", true || false)
}
