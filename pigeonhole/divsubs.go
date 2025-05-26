package pigeonhole

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// PROBLEM:
// Given a multiset of N integers, find a nonempty subset where the sum of this subset's
// elements are divisible by N
// - if one doesn't exist, return -1
// - otherwise return the size of the subset, followed by the list of elements of the multiset that comprise this subset

// Approach -> finding a subarray satisfies the problem (as a subarray is a subset)
// - Reduce to finding a subarray where the sum of its elements are divisible by N -> (Sum: S) % N == 0
// - Claim: A subarray will always exist
//   - Consider a cummlative sum of N by ranging to N and adding the next element, creating a new sum
//   - By taking the modulo of this sum S by N, we put them into buckets.
//   - Note that if the S % N == 0, we have found such a subarray that includes elements 1 to i, for i >= 1
//   - Otherwise, by the Pigeonhole Principle, at least two sums S1, S2, S2 > S1, lie in the same bucket -> S1 % N == S2 % N
//   - Therefore, N | (S2 - S1) gives us the subarray, and taking the removal of elements from S1 in S2 suffices
func divsubs(t int, nums []int) (int, []int) {
	if t <= 0 {
		return -1, nil
	}
	mod := make([]int, t)
	for i := range mod {
		mod[i] = -1
	}

	csum := 0
	var subarr []int
	for i := range nums {
		csum += nums[i]
		modsum := csum % t
		if modsum == 0 {
			for j := range i + 1 {
				subarr = append(subarr, nums[j])
			}
			return i + 1, subarr
		} else {
			if mod[modsum] == -1 {
				mod[modsum] = i
			} else {
				for j := mod[modsum] + 1; j < i+1; j++ {
					subarr = append(subarr, nums[j])
				}
				return i - mod[modsum], subarr
			}
		}
	}
	return -1, nil
}

func test_divsubs() {
	reader := bufio.NewReader(os.Stdin)

	size, _ := reader.ReadString('\n')
	size = strings.TrimSpace(size)
	n, err := strconv.Atoi(size)
	if err != nil {
		fmt.Println("Error converting size:", err)
		return
	}

	// Read the second line (slice of ints)
	ints, _ := reader.ReadString('\n')
	ints = strings.TrimSpace(ints)

	var nums []int
	for _, num := range ints {
		if num == ' ' {
			continue
		}
		tmp, err := strconv.Atoi(string(num))
		if err != nil {
			fmt.Println("Error converting number:", err)
			return
		}
		nums = append(nums, tmp)
	}

	res, subarr := divsubs(n, nums)
	fmt.Println("Size:", res, "\nSubarray:", subarr)
}
