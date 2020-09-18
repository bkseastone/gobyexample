package main

import (
	"fmt"
	"time"
)

func main() {
	size := 20000
	matrix1 := make([][]int, size)
	for i := 0; i < size; i++ {
		matrix1[i] = make([]int, size)
	}

	start1 := time.Now()
	// 方法 1
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			matrix1[i][j] = i + j
		}
	}
	fmt.Println(time.Since(start1))
	matrix2 := make([][]int, size)
	for i := 0; i < size; i++ {
		matrix2[i] = make([]int, size)
	}
	// 方法 2
	start2 := time.Now()
	// for i := 0; i < size; i++ {
	// 	for j := 0; j < size; j++ {
	// 		matrix2[j][i] = i + j
	// 	}
	// }
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			matrix2[j][i] = i + j
		}
	}
	fmt.Println(time.Since(start2))
}
