package main

import (
	"fmt"
	"time"

	sort2 "github.com/buffge/gobyexample/algorithm/sort"

	"github.com/buffge/gobyexample/rand/utils"
)

// 归并排序
func sort(arr []int, lo, hi int) {
	// 如果当前数组长度为1就返回此有序数组
	if lo < hi {
		mid := lo + hi + 1>>1
		sort(arr, lo, mid-1)
		sort(arr, mid, hi)
		merge(arr, lo, hi)
	}
}

// 合并2个有序数组
func merge(arr []int, lo, hi int) {
	// 从左右2个有序数组 以此取出最小值,当有一方为空时结束
	len := hi - lo + 1
	// 先定义有序数组
	result := make([]int, len)
	mid := lo + hi + 1>>1
	i := lo
	j := mid
	k := 0
	for i < mid && j <= hi {
		if arr[i] < arr[j] {
			result[k] = arr[i]
			k++
			i++
		} else {
			result[k] = arr[j]
			k++
			j++
		}
	}
	for i < mid {
		result[k] = arr[i]
		k++
		i++
	}
	for j <= hi {
		result[k] = arr[j]
		k++
		j++
	}
	for k, v := range result {
		arr[k] = v
	}
}
func main() {
	dataCount := sort2.DataCount
	dataCount = 100000
	arr := utils.GenerateRandomIntData(0, 100, dataCount)
	// fmt.Println(arr)
	now := time.Now()
	sort(arr, 0, len(arr)-1)
	duration := time.Now().Sub(now)
	fmt.Println(arr[0:10])
	fmt.Printf("共用时 %s\n", duration)
}
