package minimum_unsorted_subarray_to_sort_an_array

import minmax "github.com/smoothprogrammer/minmax/geeksforgeeks/minimum-and-maximum-array-elements"

func Solution(arr []int) (s, e int) {
	n := len(arr)
	if n <= 1 {
		return
	}

	// ltr scan
	element := arr[0]
	for i := 1; i < n; i++ {
		if arr[i] < element {
			s = i - 1
			break
		}
		element = arr[i]
	}

	// rtl scan
	element = arr[n-1]
	for i := n - 2; i > 0; i-- {
		if arr[i] > element {
			e = i + 1
			break
		}
		element = arr[i]
	}

	// correction
	min, max := minmax.Solution(arr[s:e])

	for i := 0; i < s; i++ {
		if arr[i] > min {
			s = i
			break
		}
	}

	for i := n - 1; i > e; i-- {
		if arr[i] < max {
			e = i
			break
		}
	}

	return s, e
}
