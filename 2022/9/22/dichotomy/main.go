package main

import (
	"fmt"
	"math/rand"
)

//从有序数组中找到num
func fromSorted(nums []int, x int) int {
	n := len(nums)
	l, r := 0, n-1
	var mid int
	for l <= r {
		mid = l + (r-l)/2
		if nums[mid] > x {
			r = mid - 1
		} else if nums[mid] < x {
			l = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

// 找到 >=num 最左的位置
func mostLeftNoLessNumIndex(nums []int, x int) int {
	n := len(nums)
	l, r := 0, n-1
	var mid int
	ans := -1
	for l <= r {
		mid = l + (r-l)/2
		if nums[mid] < x {
			l = mid + 1
		} else if nums[mid] >= x {
			ans = mid
			r = mid - 1
		}
	}
	return ans
}

// 无序，局部最小值，[0] < [1] || [n-1]<[n-2] || [i-1]<[i]<[i+1]
// 返回1个局部最小值下标
func OneMidIndex(nums []int) int {
	n := len(nums)
	if n == 0 {
		return -1
	}
	if n == 1 {
		return 0
	}
	if nums[0] < nums[1] {
		return 0
	}
	if nums[n-1] < nums[n-2] {
		return n - 1
	}

	l, r := 0, len(nums)-1
	var mid int
	ans := -1
	for l < r-1 {
		mid = l + (r-l)/2
		// 如果l,r相邻，那么mid-1就出界[l:r]
		if nums[mid] < nums[mid-1] && nums[mid] < nums[mid+1] {
			ans = mid
			break
		}
		if nums[mid] > nums[mid-1] {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return ans
}

// 生成随机数组，且相邻的数不相等
func randomArray(maxLen, maxValue int) []int {
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

// 用于测试，局部最小值
func check(nums []int, minIndex int) bool {
	if len(nums) == 0 {
		return minIndex == -1
	}
	l, r := minIndex-1, minIndex+1
	leftBigger, rightBigger := true, true
	if l >= 0 && r < len(nums) {
		leftBigger = nums[minIndex] < nums[r]
		rightBigger = nums[minIndex] < nums[l]
	}
	return leftBigger && rightBigger
}

func main() {
	// nums1 := []int{1, 2, 2, 2, 2, 3, 4, 5, 6, 6, 7, 7, 8, 9}
	// fmt.Println(fromSorted(nums1, 2))
	// fmt.Println(mostLeftNoLessNumIndex(nums1, 7))
	maxValue, maxLen := 200, 200
	fmt.Println("===测试开始===")
	for i := 0; i < 10000; i++ {
		nums := randomArray(maxLen, maxValue)
		ans := OneMidIndex(nums)
		if !check(nums, ans) {
			for _, num := range nums {
				fmt.Print(num)
				fmt.Print(",")
			}
			fmt.Println()
			fmt.Println(ans)
			break
		}
	}
	fmt.Println("===测试结束===")
}
