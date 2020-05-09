package solution

func solve(A []int) int {
	n := len(A)
	if n <= 1 {
		return n
	}
	accumulator := make(map[int]map[int]int, n)

	m := 0
	for i := 0; i < n; i++ {
		progressionLengths := make([]int, 0, n-i-1)
		for j := i + 1; j < n; j++ {
			progressionLengths = append(
				progressionLengths,
				calculate(A, n, j, A[i]-A[j], accumulator),
			)
		}
		if v := 1 + max(progressionLengths...); v > m {
			m = v
		}
	}

	return m
}

func calculate(a []int, n, index, length int, acc map[int]map[int]int) int {
	if index >= n {
		return 0
	}
	if v, ok := acc[index]; ok {
		if v2, ok := v[length]; ok {
			return v2
		}
	} else {
		acc[index] = make(map[int]int, n-index-1)
	}

	progressionLengths := make([]int, 0, n-index-1)
	for i := index + 1; i < n; i++ {
		if a[index]-a[i] == length {
			progressionLengths = append(
				progressionLengths,
				calculate(a, n, i, length, acc),
			)
		}
	}

	acc[index][length] = 1 + max(progressionLengths...)
	return acc[index][length]
}

func max(a ...int) int {
	m := 0
	for _, v := range a {
		if v > m {
			m = v
		}
	}
	return m
}
