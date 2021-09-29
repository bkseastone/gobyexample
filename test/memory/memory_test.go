package main

import (
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
