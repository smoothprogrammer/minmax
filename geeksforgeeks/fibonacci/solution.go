package fibonacci

func Solution(n int) int {
	memoization := make([]int, n+1)
	return fib(n, memoization)
}

func fib(n int, memoization []int) int {
	if n <= 1 {
		return n
	}

	if number := memoization[n]; number != 0 {
		return number
	}

	memoization[n] = fib(n-1, memoization) + fib(n-2, memoization)
	return memoization[n]
}
