package main

import (
	"fmt"
	"github.com/buffge/gobyexample/variable/sub2"
)

/**
variable
变量未初始化时
	int 为 0
	float 为 0
	布尔值为false
	字符串为""
*/

// main 包中全局变量其他包无法使用
//var Nick string
func main() {
	var i int
	var b bool
	var s string
	var f float32
	fmt.Printf("%v %v %q %v\n", i, b, s, f)
	//var i float64;
	fmt.Printf("不支持重定义变量\n")
	a := 32
	fmt.Printf("这里a会自动初始化为int = %v\n", a)
	sub2.Test()
}
