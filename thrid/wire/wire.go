// +build wireinject

package main

import "github.com/google/wire"

// 定义一个集合
var BSet = wire.NewSet(NewC, NewD, NewB)
var BSet2 = wire.NewSet(wire.Struct(new(C)), wire.Struct(new(D)), NewB)
var ASet = wire.NewSet(NewB, NewA, NewC, NewD)
var ASet2 = wire.NewSet(BSet, NewA)
var ASet3 = wire.NewSet(BSet2, NewA)

// 写法一
func InitA1() *A {
	wire.Build(NewB, NewC, NewD, NewA)
	return &A{}
}

// 写法2
func InitA2() *A {
	// 使用 Build 中 参数列表生成 函数返回值类型
	// 因为 *A 需要A B C D  4个 New函数
	wire.Build(ASet)
	return &A{}
}

func InitA3() *A {
	wire.Build(ASet2)
	return &A{}
}
func InitA() *A {
	wire.Build(ASet3)
	return &A{}
}

func InitB() *B {
	wire.Build(BSet)
	return &B{}
}
