package main

import (
	"fmt"
	"time"

	sortConf "github.com/buffge/gobyexample/algorithm/sort"
	"github.com/buffge/gobyexample/rand/utils"
)

func main() {
	dataCount := sortConf.DataCount
	// dataCount = 100_0000
	arr := utils.GenerateRandomIntData(0, 1_0000, dataCount)
	// arr = []int{1, 2, 5, 3, 1}
	// fmt.Println(arr)
	now := time.Now()
	sort(arr)
	duration := time.Now().Sub(now)
	fmt.Println(arr[0:10])
	fmt.Printf("共用时 %s\n", duration)
}

// 计数技术排序
func sort(arr []int) {
	max := 0
	for _, v := range arr {
		if v > max {
			max = v
		}
	}
	bucket := make([]int, max+1)
	for _, v := range arr {
		bucket[v]++
	}
	currIdx := 0
	for i := 0; i < max+1; i++ {
		for bucket[i] > 0 {
			arr[currIdx] = i
			currIdx++
			bucket[i]--
		}
	}
}
