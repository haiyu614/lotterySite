package utils

import (
	"lotterySite/model"
	"math/rand"


)
func GetLotteryId(goods []*model.Good) int {
    if len(goods) == 0 {
        return -1 // 边界处理：空商品列表
    }

    // 1. 计算前缀和（权重区间）
    sum := 0
    prefixSum := make([]int, len(goods))
    for i, good := range goods {
        sum += good.Number
        prefixSum[i] = sum // 前缀和：prefixSum[i] = 前i+1个商品的总数量
    }

    // 2. 生成随机数（范围：[0, sum)）
    randomIdx := rand.Intn(sum) // 比 float64 转换更简洁，且避免精度问题

    // 3. 查找随机数落在哪个前缀和区间（正确的二分逻辑）
    left, right := 0, len(prefixSum)-1
    for left < right {
        mid := (left + right) / 2
        if prefixSum[mid] <= randomIdx {
            // 随机数在 mid 之后的区间，移动左边界
            left = mid + 1
        } else {
            // 随机数可能在 mid 或之前的区间，移动右边界
            right = mid
        }
    }

    return goods[left].ID
}