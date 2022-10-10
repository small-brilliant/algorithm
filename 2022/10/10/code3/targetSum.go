package main

import "fmt"

func findTargetSumWays(nums []int, target int) int {
	sum := 0
	for _, v := range nums {
		if v < 0 {
			sum -= v
		} else {
			sum += v
		}
	}
	if (sum&1^target&1) != 0 || target > sum {
		return 0
	}
	return process2(nums, (sum+target)>>1)
}

// 从index出发，a[index:]。可以搞出rest数的个数
func process1(a []int, index, rest int, m map[int]map[int]int) int {
	if m2, ok := m[index]; ok {
		if v, ok := m2[rest]; ok {
			return v
		}
	}
	ans := 0
	if index == len(a) {
		if rest == 0 {
			ans = 1
		}
	} else {
		ans = process1(a, index+1, rest-a[index], m) + process1(a, index+1, rest+a[index], m)
	}
	if _, ok := m[index]; !ok {
		map2 := make(map[int]int, 0)
		m[index] = map2
	}
	m[index][rest] = ans
	return ans
}

func process2(a []int, target int) int {
	dp := make([]int, len(a)+1)
	dp[0] = 1
	for _, num := range a {
		for i := target; i >= num; i-- {
			dp[i] += dp[i-num]
		}
	}
	return dp[target]
}

func main() {
	// maxLen, maxV, k := 100, 100, 10
	// for i := 0; i < 100; i++ {
	// 	a := utils.RandomArray(maxLen, maxV)
	// 	ans := findTargetSumWays(a, k)

	// }
	nums := []int{1, 1, 1, 1, 1}
	fmt.Println(findTargetSumWays(nums, 3))
}
