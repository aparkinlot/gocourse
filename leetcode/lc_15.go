package main

import (
	"fmt"
	"slices"
)

// func threeSum(nums []int) [][]int {
// 	tmp := make([]int, 3)
// 	var res [][]int
// 	// [-4, -1, -1, 0, 1, 2]
// 	slices.Sort(nums)
// 	l, r, m, lc, rc := 0, len(nums)-1, 1, 2, 1
// 	tmp[0], tmp[1], tmp[2] = nums[l], nums[l+1], nums[r]
// 	//fmt.Println("HJELL")
// 	for l < r {
// 		//fmt.Println("CURR:", tmp[0], tmp[1], tmp[2])
// 		sum := tmp[0] + tmp[1] + tmp[2]
// 		if sum < 0 {
// 			if lc > rc {
// 				m = r - 1
// 				lc--
// 				rc++
// 				tmp[1] = nums[m]
// 			} else {
// 				l++
// 				tmp[0] = nums[l]
// 			}
// 		} else if sum > 0 {
// 			if lc > rc {
// 				r--
// 				tmp[2] = nums[r]
// 			} else {
// 				m = l + 1
// 				lc++
// 				rc--
// 				tmp[1] = nums[m]
// 			}
// 		} else {
// 			res = append(res, append([]int(nil), tmp...))
// 			if lc > rc {
// 				l++
// 				m = l + 1
// 				tmp[0] = tmp[1]
// 				tmp[1] = nums[m]
// 			} else {
// 				r--
// 				m = r - 1
// 				tmp[2] = tmp[1]
// 				tmp[1] = nums[m]
// 			}
// 		}
// 	}
// 	return res
// }

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

func test() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	res := threeSum(nums)
	fmt.Println(res)
	nums = []int{0, 0, 0}
	res = threeSum(nums)
	fmt.Println(res)
}
