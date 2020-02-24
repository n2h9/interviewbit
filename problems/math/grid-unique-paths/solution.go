package solution

func uniquePaths(A int, B int) int {
	if A <= 0 && B <= 0 {
		return 0
	}
	if A == 1 || B == 1 {
		return 1
	}

	// every period of time we need only upper line of field and left field value
	// also the first row and the firsl column are always eq to 1
	// row is same time previous (upper row) and currect row
	row := make([]int, B)
	// fill row with 1
	for i := 0; i < B; i++ {
		row[i] = 1
	}
	for i := 1; i < A; i++ {
		for j := 1; j < B; j++ {
			row[j] += row[j-1]
		}
	}
	return row[B-1]
}
