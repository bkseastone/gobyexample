package main

import "fmt"

func main() {
	if 7%2 == 0 {
		fmt.Println("7 是单数")
	} else {
		fmt.Println("7 是偶数")
	}

	if 8%4 == 0 {
		fmt.Println("8是4的整数倍")
	}
	//在if中可以先定义一个变量,然后整个if中都可以使用
	if num := 9; num < 0 {
		fmt.Println(num, "是负数")
	} else if num < 10 {
		fmt.Println(num, "小于10的数")
	} else {
		fmt.Println(num, "大于10的数")
	}
	//出了if就不能使用这个变量了
	//fmt.Println(num)
}
