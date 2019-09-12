package main

import (
	"fmt"
	"math"
	"time"

	sortConf "github.com/buffge/gobyexample/algorithm/sort"
	"github.com/buffge/gobyexample/rand/utils"
)

/**

 */
func main() {
	dataCount := sortConf.DataCount
	// dataCount = 10
	arr := utils.GenerateRandomIntData(-42312, 1_0000, dataCount)
	// arr = []int{1, 2, 5, 3, 1}
	// fmt.Println(arr)
	now := time.Now()
	sort(arr)
	duration := time.Now().Sub(now)
	fmt.Println(arr[0:10])
	fmt.Printf("共用时 %s\n", duration)
}

/*
 * 基数排序
 * 第一步获取绝对值最大数的位数如 12345 为5位
 * 按基数从小到大 对数组进行排序,将基数相同的放到同一个桶中
 * 遍历一轮后从桶中依次取回,此时数组相对于当前基数顺序是正确
 * 然后依次遍历下一轮基数
 */
func sort(arr []int) {
	mod := 10
	min, max := math.MaxInt64, math.MinInt64
	bucket := make([][]int, mod*2)
	for _, v := range arr {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	minAbs := min
	if min < 0 {
		minAbs = -min
	}
	if minAbs > max {
		max = minAbs
	}
	bitCount := int(math.Log(float64(max))/math.Log(float64(mod))) + 1
	var bucketIdx, currArrIdx int
	var isNegative bool
	for i := 1; i <= bitCount; i++ {
		for _, val := range arr {
			isNegative = val < 0
			base := int(math.Pow(float64(mod), float64(i)))
			radix := int(math.Abs(float64(val))) % base * mod / base
			bucketIdx = radix
			if !isNegative {
				bucketIdx += mod
			}
			bucket[bucketIdx] = append(bucket[bucketIdx], val)
		}
		currArrIdx = 0
		for j := mod - 1; j >= 0; j-- {
			for _, v := range bucket[j] {
				arr[currArrIdx] = v
				currArrIdx++
			}
			bucket[j] = bucket[j][0:0]
		}
		for j := mod; j < mod*2; j++ {
			for _, v := range bucket[j] {
				arr[currArrIdx] = v
				currArrIdx++
			}
			bucket[j] = bucket[j][0:0]
		}

	}
}
