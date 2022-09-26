package utils

import "math/rand"

// 生成随机数组，且相邻的数不相等
func RandomArray(maxLen, maxValue int) []int {
	len := int(rand.Float64() * float64(maxLen))
	nums := make([]int, len)
	if len > 0 {
		nums[0] = int(rand.Float64() * float64(maxValue))
		for i := 1; i < len; i++ {
			nums[i] = int(rand.Float64() * float64(maxValue))
			for nums[i] == nums[i-1] {
				nums[i] = int(rand.Float64() * float64(maxValue))
			}
		}
	}
	return nums
}
