package mathematicalexpectation

import (
	"fmt"
	"math"
)

// PROBLEM:
// Given an empty string S, one can perform an operation that appends a character (a-z) unofromly at random to S
// After each operation, they evaultes whether S is a palindrome or not, and if so, they get 1 point
// Suppose n operations are performed, what is the expected number of points they can get
// - the answer can be written as a/b so output a, b
func palindromes(n int) float64 {
	ev := 0.0
	for i := 1; i <= n; i++ {
		exp := float64(i / 2)
		ev += 1 / math.Pow(26.0, exp)
	}
	return ev
}

func test_palindromes() {
	fmt.Println(palindromes(3))
}
