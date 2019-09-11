package main

import (
	"fmt"
	"time"

	sortConf "github.com/buffge/gobyexample/algorithm/sort"
	"github.com/buffge/gobyexample/rand/utils"
)

func main() {
	dataCount := sortConf.DataCount
	//dataCount = 10
	arr := utils.GenerateRandomIntData(0, 100, dataCount)
	// arr = []int{1, 2, 5, 3, 1}
	// fmt.Println(arr)
	now := time.Now()
	sort(arr)
	duration := time.Now().Sub(now)
	fmt.Println(arr[0:10])
	fmt.Printf("共用时 %s\n", duration)
}

// 选择排序
// 遍历数组 寻找当前idx下最小的值.默认当前idx下最小值为 arr[minIdx](minIdx = idx)
// 再用minIdx+1- n-1 的值 依次与arr[minIdx]比较.如果后面(i)值更小,则设置minIdx = i
// 遍历完毕则替换 arr[minIdx] arr[idx]
func sort(arr []int) {
	len := len(arr)
	var minIdx int
	for i := 0; i < len-1; i++ {
		minIdx = i
		for j := i + 1; j < len; j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		if minIdx != i {
			arr[minIdx], arr[i] = arr[i], arr[minIdx]
		}
	}
}
