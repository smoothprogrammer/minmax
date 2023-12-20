package wildcard_pattern_matching

func Solution(text string, pattern string) bool {
	n, m := len(text), len(pattern)
	tStart, tEnd := 0, 0  // text window index
	pStart, pEnd := -1, 0 // pattern window index

	for tEnd < n {
		if pEnd >= m {
			break
		}

		if pattern[pEnd] == '?' || pattern[pEnd] == text[tEnd] {
			tEnd++
			pEnd++
			continue
		}

		if pattern[pEnd] == '*' {
			tStart = tEnd
			pEnd++
			pStart = pEnd
			continue
		}

		if pStart != -1 {
			tStart++
			tEnd = tStart
			pEnd = pStart
			continue
		}

		return false
	}

	for pEnd < m && pattern[pEnd] == '*' {
		pEnd++
	}

	return pEnd == m
}
