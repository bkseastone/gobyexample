package main

import (
	"fmt"
	"time"

	sortConf "github.com/buffge/gobyexample/algorithm/sort"
	"github.com/buffge/gobyexample/rand/utils"
)

func main() {
	dataCount := sortConf.DataCount
	// dataCount = 10
	arr := utils.GenerateRandomIntData(0, 100, dataCount)
	// arr = []int{1, 2, 5, 3, 1}
	// fmt.Println(arr)
	now := time.Now()
	sort(arr)
	duration := time.Now().Sub(now)
	fmt.Println(arr[0:10])
	fmt.Printf("共用时 %s\n", duration)
}

// 冒泡排序
// 从 0-n-1个元素 依次与后面的n-1 - 1 个元素进行比较如果后面的值小于前面的值则交换
// ori [1,2,5,3,1]
// 1. left:[1] right:[2,5,3,1] rightVal: 2
// ... rightVal 5,3,1
// 2. left:[1,2] right:[5,3,1] rightVal: 5
// 3. left:[1,2] right:[5,3,1] rightVal: 3
// 4. left:[1,2] right:[5,3,1] leftVal:2, rightVal: 1 -> [1,1] [5,3,2]
func sort(arr []int) {
	len := len(arr)
	tmp := 0
	for i := 0; i < len; i++ {
		for j := i + 1; j < len; j++ {
			if arr[j] < arr[i] {
				tmp = arr[i]
				arr[i] = arr[j]
				arr[j] = tmp
			}
		}
	}
}
