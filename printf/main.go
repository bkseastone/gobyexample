package main

import "fmt"

func main() {
	//%%	%字面量
	//%b	二进制整数值，基数为2，或者是一个科学记数法表示的指数为2的浮点数
	//%c	该值对应的unicode字符
	//%d	十进制数值，基数为10
	//%e	科学记数法e表示的浮点或者复数
	//%E	科学记数法E表示的浮点或者附属
	//%f	标准计数法表示的浮点或者附属
	//%o	8进制度
	//%p	十六进制表示的一个地址值
	//%s	输出字符串或字节数组
	//%T	输出值的类型，注意int32和int是两种不同的类型，编译器不会自动转换，需要类型转换。
	//%v	值的默认格式表示
	//%+v	类似%v，但输出结构体时会添加字段名
	//%#v	值的Go语法表示
	//%t	单词true或false
	//%q	该值对应的单引号括起来的go语法字符字面值，必要时会采用安全的转义表示
	//%x	表示为十六进制，使用a-f
	//%X	表示为十六进制，使用A-F
	//%U	表示为Unicode格式：U+1234，等价于"U+%04X"
	type User struct {
		Name string
		Age  int
	}
	user := User{
		"overnote",
		1,
	}
	fmt.Printf("%%\n")                   // %
	fmt.Printf("%b\n", 16)               // 10000
	fmt.Printf("%c\n", 65)               // A
	fmt.Printf("%c\n", 0x4f60)           // 你
	fmt.Printf("%U\n", '你')              // U+4f60
	fmt.Printf("%x\n", '你')              // 4f60
	fmt.Printf("%X\n", '你')              // 4F60
	fmt.Printf("%d\n", 'A')              // 65
	fmt.Printf("%t\n", 1 > 2)            // false
	fmt.Printf("%e\n", 4396.7777777)     // 4.396778e+03 默认精度6位
	fmt.Printf("%20.3e\n", 4396.7777777) //            4.397e+03 设置宽度20,精度3,宽度一般用于对齐
	fmt.Printf("%E\n", 4396.7777777)     // 4.396778E+03
	fmt.Printf("%f\n", 4396.7777777)     // 65 4396.777778
	fmt.Printf("%o\n", 16)               // 20
	fmt.Printf("%p\n", []int{1})         // 0xc000016110
	fmt.Printf("Hello %s\n", "World")    // Hello World
	fmt.Printf("Hello %q\n", "World")    // Hello "World"
	fmt.Printf("%T\n", 3.0)              // float64
	fmt.Printf("%v\n", user)             // {overnote 1}
	fmt.Printf("%+v\n", user)            // {Name:overnote Age:1}
	fmt.Printf("%#v\n", user)            // main.User{Name:"overnote", Age:1}
}
