package pigeonhole

// PROBLEM:
// Given two integers L, R where L < R, find the smallest non-negative integer N such that
// N%L > N(L+1) > ... > N%(R-1) > N%R
// - return the smallest possible N that satifies this condition
// - if not found, return -1

// Approach -> R can only be up to a value of L-1 greater than L
// - if R > 2L -1 (R >= 2L) then for any N, there will be two instances i, j where N%i == N%j
// This problem looks for all buckets to be filled (have one pigeon) for R-L+1 total buckets
// - it is clear that R%R is the lowest possible value/bucket that satisfies the condition
func decreasingSequence(l, r int) int {
	if r-l+1 > l {
		return -1
	}
	return r
}
