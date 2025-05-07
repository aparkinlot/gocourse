// A way to create state variables that can maintain a value/data structure across function calls

package fundamentals

import "fmt"

func main() {

	// anon closure
	counter := func() func(x int) int {
		current := 0
		return func(x int) int {
			fmt.Println("Pre-printed Count:", current)
			current += x
			return current
		}
	}()

	// closure -> fibonacci from Go Playground
	fibo := fib()

	for i := 0; i < 10; i++ {
		fmt.Println("Current Count:", counter(i))
		fibo()
	}
}

func fib() func() int {
	a, b := 0, 1
	return func() int {
		fmt.Println("Fibonacci:", a)
		res := a
		a, b = b, a+b
		return res
	}
}
