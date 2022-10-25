package main

import "sort"

func combinationSum(candidates []int, target int) [][]int {
	n := len(candidates)
	ans := [][]int{}
	if n == 0 {
		return ans
	}
	sort.Ints(candidates)
	path := []int{}
	sum := 0
	var backTrack func(int)
	backTrack = func(start int) {
		if sum > target || candidates[start] > target {
			return
		}
		if sum == target {
			ans = append(ans, append([]int{}, path...))
			return
		}
		for i := start; i < n; i++ {
			path = append(path, candidates[i])
			sum += candidates[i]
			backTrack(i)
			path = path[:len(path)-1]
			sum -= candidates[i]
		}
	}
	backTrack(0)
	return ans
}
