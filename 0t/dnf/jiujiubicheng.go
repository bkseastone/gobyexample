package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	N           = 1_0000_0000 // 增幅次数
	SuccessRate = 20          // 增幅成功率
)

func main() {
	// 失败次数,成功次数,当前连续失败9次的次数,非酋次数(共计连续失败n次的次数),欧皇次数(连续失败n次后 下一次成功的次数)
	var failedTimes, successTimes, consecutiveFailedTimes, feiTimes, ouTimes int
	begin := time.Now()
	rand.Seed(begin.Unix())
	for i := 0; i < N; i++ { // 增幅 N 次
		if rand.Int31n(100_00000)+1 > SuccessRate*1_00000 { // 如果增幅失败
			failedTimes++                    // 失败次数+1
			consecutiveFailedTimes++         // 当前连续失败次数+1
			if consecutiveFailedTimes == 9 { // 如果当前已经连续失败9次
				feiTimes++ // 非酋次数+1
			} else if consecutiveFailedTimes == 10 {
				consecutiveFailedTimes = 1 // 设置当前连续失败次数为1
			}
		} else { // 如果增幅成功
			successTimes++                   // 成功次数+1
			if consecutiveFailedTimes == 9 { // // 如果当前已经连续失败9次
				ouTimes++ // 欧皇次数加1
			}
			consecutiveFailedTimes = 0 // 重置当前连续失败次数
		}
	}
	fmt.Printf(`
共运行 %d 次 耗时%.2fs
当前成功几率: %.2f
成功 %d 次,失败 %d 次
非酋次数: %d 次,欧皇次数: %d 次,欧皇概率 %.2f
`[1:], N, time.Since(begin).Seconds(), SuccessRate/100.0, successTimes,
		failedTimes, feiTimes, ouTimes, float64(ouTimes)/float64(feiTimes))
}
