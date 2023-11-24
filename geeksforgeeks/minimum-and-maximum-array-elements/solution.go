package minimum_and_maximum_array_elements

func Solution(arr []int) (min, max int) {
	length := len(arr)

	if length == 0 {
		return 0, 0
	}

	if length == 1 {
		return arr[0], arr[0]
	}

	min, max = arr[0], arr[1]
	if min > max {
		min, max = max, min
	}

	for i := 2; i < length; i++ {
		v := arr[i]

		if v < min {
			min = v
		}

		if v > max {
			max = v
		}
	}

	return min, max
}
