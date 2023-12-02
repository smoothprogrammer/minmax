package chocolate_distribution_problem

import "slices"

func Solution(arr []int, m int) (mindiff int) {
	n := len(arr)
	if n <= 1 {
		return 0
	}

	slices.Sort(arr)

	mindiff = arr[n-1]
	for i := 0; i < n-m; i++ {
		diff := arr[i+m-1] - arr[i]
		if diff < mindiff {
			mindiff = diff
		}
	}

	return mindiff
}
