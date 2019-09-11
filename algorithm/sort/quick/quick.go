package main

import (
	"fmt"
	"time"

	sort2 "github.com/buffge/gobyexample/algorithm/sort"

	"github.com/buffge/gobyexample/rand/utils"
)

// 归并排序
func sort(arr []int, lo, hi int) {
	if lo < hi {
		mid := split(arr, lo, hi)
		sort(arr, lo, mid-1)
		sort(arr, mid+1, hi)
	}
}

// 计算出分割idx
func split(arr []int, lo, hi int) int {
	pivot := lo
	mid := pivot + 1
	for i := mid; i <= hi; i++ {
		if arr[i] < arr[pivot] {
			arr[i], arr[mid] = arr[mid], arr[i]
			mid++
		}
	}
	arr[pivot], arr[mid-1] = arr[mid-1], arr[pivot]
	return mid - 1
}
func main() {
	dataCount := sort2.DataCount
	dataCount = 100_0000
	arr := utils.GenerateRandomIntData(0, 1_0000, dataCount)
	// fmt.Println(arr)
	now := time.Now()
	sort(arr, 0, len(arr)-1)
	duration := time.Now().Sub(now)
	fmt.Println(arr[0:10])
	fmt.Printf("共用时 %s\n", duration)
}
