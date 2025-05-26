package pigeonhole

// PROBLEM: (Unfinished)
// Given N 64-bit integers such that any two successive numbers (i, i+1) differ only by one bit,
// find out if there exists four integers in the array such that their XOR is equal to 0
// - return yes if there is, no if not

// Approach -> Consider an array that takes the XOR of pairs of two successive numbers (<i, i+i>, <i+2, i+3>, ... , <i+N-2, i+N-1>)
// - note that because these two numbers are successive (and thus differ by a bit), we can formulate buckets
// - that is 64 buckets for 64 bit positions
//   - for N >= 130, that means, this approach will put two pairs into the same bucket and thus, the XOR of these
//	   two pairs is 0
// The approach only works for N >= 130
// - otherwise, we default to a brute force that takes the XOR of three numbers in N: a ^ b ^ c
// - find if there extists a index d such that d = a ^ b ^ c, and of course, these indices must all be unique in N
func gsc(nums []int64) string {
	return ""
}

func test_gsc() {

}
