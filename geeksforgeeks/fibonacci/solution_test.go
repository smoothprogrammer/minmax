package fibonacci_test

import (
	"strconv"
	"testing"

	"github.com/minizilla/minmax/geeksforgeeks/fibonacci"
	"github.com/minizilla/testr"
)

func TestFibonacci(t *testing.T) {
	fib := []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55}

	for n, number := range fib {
		t.Run(strconv.Itoa(n), func(t *testing.T) {
			assert := testr.New(t)
			out := fibonacci.Solution(n)
			assert.Equal(out, number)
		})
	}
}
