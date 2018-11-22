package main

import (
	"flag"
	"fmt"
)

/**
  ./main -str 这是一个字符串 -int 4396 -bool -svar 随便写个字符串  sdf 234 sdf
*/
func main() {
	//定义一个选项 默认值 值的介绍
	wordPtr := flag.String("str", "foo", "a string")
	numbPtr := flag.Int("int", 42, "an int")
	//如果是bool值 只需要 -bool 就行  不要 = true
	boolPtr := flag.Bool("bool", false, "a bool")
	var svar string
	//将字符串保存到变量中
	flag.StringVar(&svar, "svar", "bar", "a string var")
	flag.Parse()
	fmt.Println("str:", *wordPtr)
	fmt.Println("int:", *numbPtr)
	fmt.Println("bool:", *boolPtr)
	fmt.Println("svar:", svar)
	fmt.Println("尾部的其他参数 tail:", flag.Args())
}
