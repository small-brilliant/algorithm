package main

func nextPermutation(nums []int) {
	n := len(nums)
	i := n - 1
	for i > 0 && nums[i] <= nums[i-1] {
		i--
	}
	if i == 0 {
		resver(nums, 0, n-1)
	} else {
		for j := n - 1; j >= i-1; j-- {
			if nums[j] > nums[i-1] {
				nums[i-1], nums[j] = nums[j], nums[i-1]
				resver(nums, i, n-1)
				return
			}
		}
	}
}
func resver(nums []int, i, j int) {
	for l, r := i, j; l <= r; l, r = l+1, r-1 {
		nums[r], nums[l] = nums[l], nums[r]
	}
}

func main() {
	nextPermutation([]int{2, 3, 1})
}
