package main

func searchRange(nums []int, target int) []int {
	n := len(nums)
	l, r := 0, n
	mid := l + (r-l)>>1
	for l < r {
		if nums[mid] == target {
			break
		}
		if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid + 1
		}
		mid = l + (r-l)>>1
	}
	if nums[mid] != target {
		return []int{}
	}
	i, j := mid, mid
	for ; i < n; i++ {
		if nums[i] == target {
			continue
		}
	}
	if i == n {
		i--
	}
	for ; i >= 0; j-- {
		if nums[i] == target {
			continue
		}
	}
	if i == -1 {
		i++
	}
	return []int{i, j}
}

func main() {
	searchRange([]int{1, 2, 3, 4, 4, 4, 4, 5}, 4)
}
