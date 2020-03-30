package main

import (
	"fmt"
)

func main() {
	// 顶一个长度为3的string类型切片
	s := make([]string, 0, 1)
	fmt.Printf("切片s:%v,长度为%d,容量为%d\n", s, len(s), cap(s))
	s = append(s, "0")
	s = append(s, "1")
	s = append(s, "2")
	s[0] = "0"
	s[1] = "1"
	s[2] = "2"
	fmt.Println("现在s的值为: ", s)
	fmt.Println("s[2]: ", s[2])

	fmt.Println("s的长度为: ", len(s))
	// 想切片中追加
	s = append(s, "3")
	s = append(s, "4", "5")
	fmt.Println("添加之后 s: ", s)
	// 定义c为长度为len(s)的字符串切片
	c := make([]string, len(s))
	// 复制s到c上
	copy(c, s)
	fmt.Println("复制s到c之后 c: ", c)
	l := s[2:5]
	fmt.Println("l是s的第[2-5)区间的切片,现在l: ", l)
	fmt.Printf("l的类型是%T\n", l)

	l = s[:5]
	fmt.Println("l是s的第[begin-5)区间的切片,现在l: ", l)

	l = s[2:]
	fmt.Println("l是s的第[2-end]区间的切片,现在l: ", l)
	// 不设置长度的切片,自动设置长度
	t := []string{"6", "7", "8"}
	t = append(t, "9")
	fmt.Println("添加完j之后 t: ", t)

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("二维切片: ", twoD)
}
