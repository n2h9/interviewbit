package solution

func prettyPrint(A int) [][]int {
	if A <= 0 {
		return nil
	}

	if A == 1 {
		return [][]int{{1}}
	}

	n := A*2 - 1
	m := make([][]int, n)
	for i := 0; i < n; i++ {
		m[i] = make([]int, n)
	}
	min := 0
	max := n - 1
	for A > 0 {
		for i := min; i <= max; i++ {
			m[min][i] = A
			m[max][i] = A
			m[i][min] = A
			m[i][max] = A
		}
		min++
		max--
		A--
	}

	return m
}
