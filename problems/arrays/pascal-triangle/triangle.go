package triangle

func solve(n int) [][]int {
	if n <= 0 {
		return nil
	}
	a := make([][]int, n)
	a[0] = []int{1}

	for i := 1; i < n; i++ {
		a[i] = make([]int, i+1)
		a[i][0] = 1
		for j := 1; j <= i; j++ {
			a[i][j] = a[i-1][j-1]
			if j < i {
				a[i][j] += a[i-1][j]
			}
		}
	}

	return a
}
