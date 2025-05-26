package leetcode

import (
	"fmt"
	"slices"
)

func threeSum(nums []int) [][]int {
	var res [][]int
	slices.Sort(nums)

	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue // skip duplicate 'i'
		}

		left, right := i+1, len(nums)-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum < 0 {
				left++
			} else if sum > 0 {
				right--
			} else {
				res = append(res, []int{nums[i], nums[left], nums[right]})

				// Skip duplicate 'left' and 'right'
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}

				left++
				right--
			}
		}
	}
	return res
}

func test_threeSum() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	res := threeSum(nums)
	fmt.Println(res)
	nums = []int{0, 0, 0}
	res = threeSum(nums)
	fmt.Println(res)
}
