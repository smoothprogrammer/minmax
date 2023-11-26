package minimum_swaps_to_sort_an_array

func Solution(arr []int) (swaps int) {
	n := len(arr)
	if n <= 1 {
		return 0
	}

	var opening int
	cycle := make([]int, n)
	for i := 0; i < n; i++ {
		icurr := i
		v := arr[icurr]
		inext := v - 1

		if arr[inext] == cycle[inext] &&
			cycle[icurr] > cycle[inext] {
			opening--
		}

		if cycle[icurr] < cycle[inext] {
			cycle[icurr] = cycle[inext]
		} else {
			cycle[inext] = cycle[icurr]
		}

		if cycle[icurr] == 0 {
			cycle[icurr] = v
			opening++
		}

		if cycle[inext] == 0 {
			cycle[inext] = v
		}
	}

	return n - opening
}
