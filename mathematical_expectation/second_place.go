package main

import (
	"cmp"
	"fmt"
	"slices"
)

// PROBLEM:
// There are n players participating in a knockout tournament. The i-th player has the rating a_i
// - No two players have the same rating
// A total of n-1 matches are played in the tournament.
// - In each match, 2 random players are chosen of the remaining players and the higher rated one wins, loser is out
// What is the expected rating of the player ending up at second place (second in the tournament)
// - The poor fellow loses the final match

// First compute the expected rating of the second-place player (the runner-up) in a random knockout tournament where:
// - Players are eliminated one-by-one in random pairings -> thus our pairings are i+1 Choose 2.
// - We know dhe player with the higher rating always wins.
// - The strongest player will always win the tournament, so the runner-up is the last player they beat.
// Each new player gets a chance to be runner-up
// Blend that in using their probability and the previous answer — and let linearity of expectation do the rest.

// At each step: "What's the chance this newly added player ends up facing the strongest player in the final match (and loses)?"
// That chance is: 1 / i+1 Choose 2 -> 1 / (i*(i+1))/2
// - Simplified to 2 / (i*(i+1))
// - that pair being <strongest_player, player_i> across all possible pairs up to and including i+1
//   - ie: player 0 (strongest), player 1, player 2 -> i=0, i=1, i=2 -> three players -> i+1 (3)

// Linearity of Expectation -> this allows us to blend in a new possible outcome with known probability to existing expected value group
// -> (1-P(i))*(Previous EV) + P(i)*a[i]
//    - we can do this as we are computing a weighted sum and thus, we are simply adjusting it (previous EV) to factor in the new probability

func secondPlace(n int, a []int) float64 {
	cmpFn := func(a, b int) int {
		return cmp.Compare(b, a)
	}
	slices.SortFunc(a, cmpFn)
	ans := make([]float64, n)
	if n > 1 {
		ans[1] = float64(a[1])
	} else {
		return -1
	}

	for i := 2; i < n; i++ {
		val := float64(i*(i+1)) / 2.0
		inv := 1.0 / val

		ans[i] = (1.0-inv)*ans[i-1] + inv*float64(a[i])
	}

	return ans[n-1]
}

func main() {
	a := []int{1, 2, 3}
	fmt.Println(secondPlace(len(a), a))
}
