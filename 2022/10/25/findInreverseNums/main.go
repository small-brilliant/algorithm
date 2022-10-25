package main

func search(nums []int, target int) int {
	n := len(nums)
	l, r := 0, n-1
	mid := l + (r-l)>>1
	for l <= r {
		if nums[mid] == target {
			return mid
		}
		if nums[mid] >= nums[0] {
			if target < nums[mid] && target >= nums[0] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else {
			if target > nums[mid] && target <= nums[n-1] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
		mid = l + (r-l)>>1
	}
	return -1
}
func main() {
	search([]int{4, 5, 6, 7, 0, 1, 2}, 0)
}
