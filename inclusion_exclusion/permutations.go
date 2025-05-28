package main

import (
	"fmt"
)

func permutations() int {
	univ := 10 * 9 * 8 * 7 * 6 * 5 * 4 * 3 * 2

	// singles -> (13, 49, 34, 23)
	cnt := 4 * (9 * 8 * 7 * 6 * 5 * 4 * 3 * 2)
	// pairs -> (134, 234, 349, 13&49, 49&23)
	cnt = cnt - (3 * (8 * 7 * 6 * 5 * 4 * 3 * 2)) - (2 * (7 * 6 * 5 * 4 * 3 * 2))
	// triples -> (1349, 2349)
	cnt = cnt + (2 * (7 * 6 * 5 * 4 * 3 * 2 * 1))
	// no quadruples

	return univ - cnt
}

func main() {
	fmt.Println("Perms:", permutations())
}
