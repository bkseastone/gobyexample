package main

import "fmt"

// 归并排序
func mergeSort(arr []int) []int {
	arrLen := len(arr)
	// 如果当前数组长度为1就返回此有序数组
	if arrLen == 1 {
		return arr
	}
	// 数组一分为二
	mid := arrLen >> 1
	left := arr[0:mid]
	right := arr[mid:]
	// 对分开的2个数组分别排序
	// 返回合并后的有序数组
	return merge(mergeSort(left), mergeSort(right))
}

// 合并2个有序数组
func merge(left []int, right []int) []int {
	// 先定义有序数组
	var result []int
	// 从左右2个有序数组 以此取出最小值,当有一方为空时结束
	for len(left) != 0 && len(right) != 0 {
		if left[0] <= right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}
	// 如果取最小值结束时 左数组还有数值,那么左数组中值就是最大值,
	// 直接加到尾巴上
	if len(left) != 0 {
		result = append(result, left...)
	}
	if len(right) != 0 {
		result = append(result, right...)
	}
	return result
}
func main() {
	arr := []int{35, 234, 1, 461, 6, 12, 61, 3235, 436, 12, 3}
	fmt.Println(mergeSort(arr))
}
