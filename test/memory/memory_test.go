package main

import (
	"reflect"
	"testing"
)

// 发生扩容时会分配内存
func BenchmarkMem1(b *testing.B) {
	arr := [5]int{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// 此变量在栈上
		tmpArr := make([]int, len(arr))
		tmpArr = append(tmpArr, arr[:]...) // 扩容一倍 分配 5*2 * 8  = 80b
		tmpArr = append(tmpArr, i)         // 扩容一倍 10*2 *8 = 160b
		// 共分配2次 240b
	}
}
func fnReturnInterface1() interface{} {
	res := int32(255)
	return res
}
func fnReturnInterface2() interface{} {
	res := 256
	return res
}

// 返回interface{}会分配内存
func BenchmarkMem2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fnReturnInterface1() // 0-256的小整数 不分配内存
		fnReturnInterface2() // 分配一次 8个字节(64位机器int)
	}
}
func fnReflectGetType1() string {
	var res interface{} = 255
	t := reflect.TypeOf(res).Kind()
	return t.String()
}
func fnReflectGetType2(reflectA interface{}) string {
	t := reflect.TypeOf(reflectA).Kind()
	return t.String()
}

type (
	userI interface {
		Age() int
	}
	userS struct { // 32字节
		age  int
		age2 int
		age3 int32
		age4 int16
		age5 int8
		age6 int8
		age7 int8 // 字节对齐的问题 如果去掉7 就只要24字节了
	}
)

func (u userS) Age() int {
	return u.age
}

// 在interface类型上调用方法会分配内存
func BenchmarkMem3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var u userI
		u = userS{age: 1232}
		u.Age()         // interface调用 分配32字节内存(实际struct内存)
		u.(userS).Age() // 实际类型调用
	}
}
func fnTestChanMem1() {
	c := make(chan int, 0)     // 1 alloc 96字节 runtime.hchan 结构体占96字节
	c2 := make(chan string, 1) // 2 alloc 102字节 runtime.hchan 结构体占96字节 1个string结构体占16字节
	c3 := make(chan string, 5) // 2 alloc 176字节 runtime.hchan 结构体占96字节 5个string结构体占80字节
	// c <- a
	// _ = a
	_ = c
	_ = c2
	_ = c3
}

// 创建chan时会在堆上申请(网上没看到,自己测出来)
func BenchmarkMem4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// a := new(userS)
		fnTestChanMem1()
		// _ = <-c
	}
}

// 向chan传递指针时会导致 变量逃逸到堆
func BenchmarkMem5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a := ""
		c := make(chan *string, 1) // 2 alloc 104字节 runtime.hchan 结构体占96字节 一个指针 占8字节
		c <- &a                    // 向chan传递指针时会导致 变量逃逸到堆 // 1 alloc 16字节 string
		_ = c
	}
}

// 向chan传递带指针的结构体会导致 变量逃逸到堆
func BenchmarkMem6(b *testing.B) {
	type s1 struct {
		age   *int8
		name  string
		name2 string
	}
	for i := 0; i < b.N; i++ {
		age := new(int8)
		a := s1{
			age: age,
		}
		c := make(chan s1, 1) // 2 alloc 144字节 runtime.hchan 结构体占96字节 s1结构体 分配一次 40字节 但有字节对齐
		// 从5*8 -> 6*8 即48字节
		c <- a // 向chan传递带指针的结构体会导致 指针变量逃逸到堆 // 1 alloc 1个字节 int8
	}
}

// 在一个切片上存储指针或带指针的值
func BenchmarkMem7(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a := make([]*int16, 0, 10)
		t := new(int16) // 1 alloc 2字节 new(int16)发生逃逸
		a = append(a, t)
		a2 := make([]*userS, 0, 10)
		t2 := new(userS) // 1 alloc 32字节   new(userS) 逃逸
		a2 = append(a2, t2)
	}
}

// 无法在编译时确定的大小
func BenchmarkMem8(b *testing.B) {
	fn := func(n int) {
		_ = make([]int8, 0, n)
	}
	n := 5
	const n2 = 5
	for i := 0; i < b.N; i++ {
		fn(n)                   // 1 alloc  5个字节
		_ = make([]int8, 0, n)  // 1 alloc 5个字节
		_ = make([]int8, 0, n2) // 0 alloc 常量在编译时就知道大小 所以会在栈上
	}
}

// todo 协程执行函数会分配
func BenchmarkMem9(b *testing.B) {
	nop := func() {}
	fn1 := func() int8 {
		return 0
	}
	_ = nop
	for i := 0; i < b.N; i++ {
		go fn1() // 1 alloc 16b todo why 16b
		go nop()
		go fn1() // 1 alloc 16b
	}
}
