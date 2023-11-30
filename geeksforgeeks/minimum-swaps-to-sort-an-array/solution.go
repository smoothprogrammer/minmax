package minimum_swaps_to_sort_an_array

func Solution(arr []int) (swaps int) {
	n := len(arr)
	if n <= 1 {
		return 0
	}

	var cycle int
	opening := make([]int, n)
	for i := 0; i < n; i++ {
		icurr := i
		v := arr[icurr]
		inext := v - 1

		if arr[inext] == opening[inext] &&
			opening[icurr] > opening[inext] {
			cycle--
		}

		if opening[icurr] < opening[inext] {
			opening[icurr] = opening[inext]
		} else {
			opening[inext] = opening[icurr]
		}

		if opening[icurr] == 0 {
			opening[icurr] = v
			cycle++
		}

		if opening[inext] == 0 {
			opening[inext] = v
		}
	}

	return n - cycle
}
